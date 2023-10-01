use csv::Reader;
use economy_core::Ohlc;

use super::OhlcCsvParser;


pub struct OhlcCsvParserFile {}

impl OhlcCsvParser for OhlcCsvParserFile {
    fn parse(mut reader: Reader<&[u8]>) -> Vec<Ohlc> {
        // Detect the header line
        // Based on the header, create the column
        // for record in reader.records() {
        //     let record = record?;
        //     println!(
        //         "In {}, {} built the {} model. It is a {}.",
        //         &record[0],
        //         &record[1],
        //         &record[2],
        //         &record[3]
        //     );
        // }

        Vec::new()
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn detect_header() {
        let header1 = "<TICKER>,<DTYYYYMMDD>,<TIME>,<OPEN>,<HIGH>,<LOW>,<CLOSE>,<VOL>";
        let mut reader = csv::Reader::from_reader(header1.as_bytes());
    }
}