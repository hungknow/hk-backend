use std::{path::Path, fs::File};

use csv::Reader;

use crate::{Ohlc, MarketGenerator, Result};

/// The lifetime parameter `'r` refers to the lifetime of the underlying
/// CSV `Reader`.
pub struct OhlcCsvParserFile<'r, R: 'r> {
    // the reader contains the parsed row
    // reader : Reader<&[u8]>,
    reader: &'r mut R,
}

// impl OhlcCsvParser for OhlcCsvParserFile {
//     fn parse(mut reader: Reader<&[u8]>) -> Option<Result<Vec<Ohlc>, EconomyError>> {
//         // Detect the header line
//         // Based on the header, create the column
//         for record in reader.records() {
//             match record {
//                 Ok(record) => Some(record),
//                 Err(err) => Err(err),
//             }
//         }
//         None
//     }
// }

impl<'r, R> OhlcCsvParserFile<'r, csv::Reader<R>> {
    pub fn from_reader(reader: &'r mut csv::Reader<R>) -> Result<Self> {
        Ok(OhlcCsvParserFile {
            reader: reader,
        })
    }

    // pub fn from_path<P: AsRef<Path>>(file_path: P) -> Result<Self> {
    //     let mut reader = Reader::from_path(file_path)?;
    //     // match reader {
    //     //     Ok(mut reader) => Ok(OhlcCsvParserFile {
    //     //         reader: &mut reader,
    //     //     }),
    //     //     Err(err) => Err(err.into()),
    //     // }
    //     Ok(OhlcCsvParserFile {
    //         reader: &reader,
    //     })
    
    // }
}


// impl MarketGenerator<Ohlc> for OhlcCsvParserFile {
//     fn next(&mut self) -> crate::Feed<Ohlc> {
//         for record in self.reader.records() {
//             record
//         }
//         todo!()
//     }
// }

#[cfg(test)]
mod tests {
    use crate::OhlcCsvParserFile;

    #[test]
    fn detect_header() {
        let header1 = "<TICKER>,<DTYYYYMMDD>,<TIME>,<OPEN>,<HIGH>,<LOW>,<CLOSE>,<VOL>";
        let mut reader = csv::Reader::from_reader(header1.as_bytes());
    }

    // Read content from file
    // Iterate over the content until the end of file
    #[test]
    fn new() {
        let data = "\
     city,country,popcount
     Boston,United States,4628910
     ";
        let mut reader = csv::Reader::from_reader(data.as_bytes());
        let mut ohlcCsvParser = OhlcCsvParserFile::from_reader(&mut reader);
        // ohlcCsvParser.next();
    }
}