package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/petuhovskiy/cian-search/jsv"
	"github.com/petuhovskiy/cian-search/nometa"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var tokenFile string

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := tokenFile
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

var sheetID string   // long nonsense doc id string
var sheetPage string // Sheet1

func letter(i int) string {
	const cnt = 'Z' - 'A' + 1
	if i > cnt {
		return letter((i-1)/cnt) + letter((i-1)%cnt+1)
	}
	return string(rune(i + 'A' - 1))
}

func rangeOf(r string, args ...interface{}) string {
	if len(args) == 0 {
		return sheetPage + "!" + r
	}

	return sheetPage + "!" + fmt.Sprintf(r, args...)
}

func fixHeader(cli *sheets.Service, sample map[string]string) (map[string]int, error) {
	readRange := rangeOf("A1:1")
	resp, err := cli.Spreadsheets.Values.Get(sheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %w", err)
	}
	spew.Dump(resp)

	mx := 0

	res := make(map[string]int)
	if len(resp.Values) > 0 {
		for i, val := range resp.Values[0] {
			res[fmt.Sprintf("%v", val)] = i + 1
			if i > mx {
				mx = i + 1
			}
		}
	}

	for key := range sample {
		if _, ok := res[key]; ok {
			continue
		}

		mx++
		res[key] = mx

		value := &sheets.ValueRange{
			Values: [][]interface{}{{key}},
		}

		rng := rangeOf("%v1:%v1", letter(mx), letter(mx))
		spew.Dump(rng)

		wResp, err := cli.Spreadsheets.Values.
			Update(sheetID, rng, value).
			ValueInputOption("USER_ENTERED").
			Do()
		spew.Dump(wResp)
		if err != nil {
			return nil, fmt.Errorf("failed to update header at %v: %w", mx, err)
		}
	}

	return res, nil
}

func valueGet(pos int, values [][]interface{}) string {
	if len(values) == 0 {
		return ""
	}
	pos = pos - 1

	if len(values[0]) > pos {
		return fmt.Sprintf("%v", values[0][pos])
	}
	return ""
}

func colorRowRed(cli *sheets.Service, row int) error {
	req := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			&sheets.Request{
				RepeatCell: &sheets.RepeatCellRequest{
					Range: &sheets.GridRange{
						// SheetId:       sheetID,
						StartRowIndex: int64(row - 1),
						EndRowIndex:   int64(row),
					},
					Cell: &sheets.CellData{
						UserEnteredFormat: &sheets.CellFormat{
							BackgroundColor: &sheets.Color{
								Blue:  0.5,
								Green: 0.5,
								Red:   1.0,
							},
						},
					},
					Fields: "userEnteredFormat(backgroundColor)",
					// "fields": "userEnteredFormat(backgroundColor,textFormat,horizontalAlignment)"
				},
			},
		},
	}

	_, err := cli.Spreadsheets.BatchUpdate(sheetID, req).Do()
	if err != nil {
		return err
	}

	return nil
}

func writeValue(cli *sheets.Service, row int, column int, val interface{}) error {
	value := &sheets.ValueRange{
		Values: [][]interface{}{{val}},
	}

	rng := rangeOf("%v%v", letter(column), row)
	_, err := cli.Spreadsheets.Values.
		Update(sheetID, rng, value).
		ValueInputOption("USER_ENTERED").
		Do()
	return err
}

type sheetWrite struct {
	row int
	col int
	val interface{}
}

func writeManyInRow(cli *sheets.Service, arr []sheetWrite) error {
	mn := 1
	mx := 1
	for _, v := range arr {
		if v.col > mx {
			mx = v.col
		}

		if v.row != arr[0].row {
			return fmt.Errorf("mismatching row, %v and %v", v.row, arr[0].row)
		}
	}

	vals := make([]interface{}, mx)
	for _, v := range arr {
		vals[v.col-1] = v.val
	}

	value := &sheets.ValueRange{
		Values: [][]interface{}{vals},
	}

	rng := rangeOf("%v%v:%v%v", letter(mn), arr[0].row, letter(mx), arr[0].row)
	_, err := cli.Spreadsheets.Values.
		Update(sheetID, rng, value).
		ValueInputOption("USER_ENTERED").
		Do()
	return err
}

func updateRow(cli *sheets.Service, row int, header map[string]int, offer nometa.Offer) error {
	// sleeping before update row
	time.Sleep(time.Second)

	cur, err := jsv.Marshal(offer)
	if err != nil {
		return err
	}

	var updates []sheetWrite
	for k, v := range cur {
		pos, ok := header[k]
		if !ok {
			return fmt.Errorf("header %v not found", k)
		}

		updates = append(updates, sheetWrite{
			row: row,
			col: pos,
			val: v,
		})
	}

	return writeManyInRow(cli, updates)
}

func main() {
	var credsFile string
	var resultFile string
	var startRow int

	flag.StringVar(&credsFile, "creds", "credentials.json", "path to google creds.json file")
	flag.StringVar(&sheetID, "sheet", "value_here", "google sheet doc id")
	flag.StringVar(&tokenFile, "tokenfile", "token.json", "path to save oauth token")
	flag.StringVar(&resultFile, "result", "result.json", "path with cian results")
	flag.StringVar(&sheetPage, "page", "Sheet1", "google sheets bottom page")
	flag.IntVar(&startRow, "startRow", 2, "row to start iteration from")
	flag.Parse()

	b, err := ioutil.ReadFile(credsFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	cli, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	jsonData, err := ioutil.ReadFile(resultFile)
	if err != nil {
		log.Fatalf("Failed to read result file: %v", err)
	}

	var results []nometa.Offer
	err = json.Unmarshal(jsonData, &results)
	if err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)
	}

	if len(results) == 0 {
		panic("results is empty")
	}

	resultsMap := make(map[string]nometa.Offer)
	for _, offer := range results {
		resultsMap[strconv.Itoa(offer.CianID)] = offer
	}

	sample, err := jsv.Marshal(results[0])
	if err != nil {
		panic(err)
	}

	header, err := fixHeader(cli, sample)
	if err != nil {
		panic(err)
	}
	spew.Dump(header)

	fromTable := make(map[string]int)

	mxLine := 1
	for i := startRow; ; i++ {
		log.Printf("Progress: line %d", i)

		readRange := rangeOf("A%d:%d", i, i)
		resp, err := cli.Spreadsheets.Values.Get(sheetID, readRange).Do()
		if err != nil {
			panic(err)
		}

		if valueGet(header["CianID"], resp.Values) == "" {
			break
		}
		mxLine = i

		currentCianID := valueGet(header["CianID"], resp.Values)
		currentCianID = strings.ReplaceAll(currentCianID, ",", "")
		fromTable[currentCianID] = i

		if valueGet(header["IsDeleted"], resp.Values) == "TRUE" {
			time.Sleep(time.Second)
			continue
		}

		offer, ok := resultsMap[currentCianID]
		if !ok {
			err := colorRowRed(cli, i)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second)

			offer.IsDeleted = true
			err = writeValue(cli, i, header["IsDeleted"], "TRUE")
			if err != nil {
				panic(err)
			}
			continue
		}
		err = updateRow(cli, i, header, offer)
		if err != nil {
			panic(err)
		}
	}

	for _, offer := range results {
		if _, ok := fromTable[strconv.Itoa(offer.CianID)]; ok {
			continue
		}

		mxLine = mxLine + 1
		log.Printf("Progress: line %d", mxLine)

		err := updateRow(cli, mxLine, header, offer)
		if err != nil {
			panic(err)
		}
	}
}
