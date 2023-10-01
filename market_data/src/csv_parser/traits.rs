use economy_core::Ohlc;

pub trait OhlcCsvParser {
    fn parse(reader: csv::Reader<&[u8]>) -> Vec<Ohlc>;
}
