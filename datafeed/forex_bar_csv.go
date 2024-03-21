package datafeed

import (
	"io/fs"

	"github.com/spf13/afero"
	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

// XAUUSD_M1 => CSVSymbolFiles
type ResolutionToSymbolToCsvCandleFile map[string]*CSVSymbolFiles

// Keep the list of CSV and the symbol info
// Read the CSV and produce the bars
type ForexBarCsvManager struct {
	ResolutionToSymbolToCsvCandleFile ResolutionToSymbolToCsvCandleFile
}

type FromTimeToCSVCandleFile map[int64]*CSVCandleFile

type CSVCandleFileManager interface {
	GetCSVFilesByTimeRange(ticket string, resolution models.Resolution, fromTimestamp int64, toTimeStamp int64) (FromTimeToCSVCandleFile, error)
	GetCandlesFromCsvFile(csvCandleFile *CSVCandleFile, fromTimestamp int64, toTimestamp int64) (*models.Candles, error)
}

/*
*
Read CSV from 1M CSV.
*/
type CSVCandleFile struct {
	Path      string
	TimeRange models.TimeRange
}

type CSVSymbolFiles struct {
	SymbolInfo models.SymbolInfo
	Resolution models.Resolution

	// Map from the from timestamp to CSVBarFile
	Files map[int64]*CSVCandleFile
}

type ForexBarCsvProducer interface {
	// GetBars reads the CSV file and returns the bars.
	// Produce 5M, 15M, 1H bars from 1M CSV files.
}

func (o *ForexBarCsvManager) ScanFolder(
	folderPath string,
) error {
	// The name of folder is symbol name
	// Then the folder name is resolution
	// Then the file name is the timestamp: XAUUSD_2023_06.csv
	return nil
}

// Root dir:
// - M1
//   - XAUUSD_2023_06.csv
//   - XAUUSD_2023_07.csv
//
// - M5
//   - XAUUSD_2023_06.csv
//   - XAUUSD_2023_07.csv
func scanRootDir() {

}

func scanFolder(fsAPI afero.Fs, folderPath string) ([]string, error) {
	err := afero.Walk(fsAPI, folderPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {

		}
		// Read the file name and create a CSVBarFile
		return nil
	})
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppInternalError, "Failed to scan the folder %s: %s", folderPath, err.Error())
	}
	// Read the folder and return the list of file names
	return nil, nil
}

func (o *ForexBarCsvManager) findCSVFilesByTimeRange(ticket string, resolution models.Resolution, fromTimestamp int64, toTimeStamp int64) (*CSVCandleFile, error) {
	// Get the array of files by ticket and resolution
	// Find the files that are in the time range
	return nil, nil
}

func (o *ForexBarCsvManager) GetBars(
	symbolInfo *models.SymbolInfo,
	resolution models.Resolution,
	periodParams *models.PeriodParams,
) (*models.GetBarsResult, error) {
	// ticket := symbolInfo.Ticket
	// resolutionStr := resolution.String()

	// Create the time block from the input time range and the resolution
	// timeBlocks := timeutils.CreateTimeBlock(periodParams.FromTimestamp, periodParams.ToTimestamp, resolution.Seconds())
	// Find all CSV files by time blocks
	// Query each CSV file and read the bars
	// Composite all bars and return
	panic("Not implemented")
}
