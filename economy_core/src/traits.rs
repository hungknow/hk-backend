use crate::{Ohlc, EconomyError};

pub trait OhlcCsvParser {
    fn parse(reader: csv::Reader<&[u8]>) -> Vec<Ohlc>;
}


pub trait OhlcReader {
    fn get(&self, symbol: &str, query: &str) -> Vec<Ohlc>;
}

#[derive(Clone, Copy, Debug)]
pub struct OhlcReplayState {
    pub read_index: u32,
    pub available_count: u32,
}

/**
 * Activity:
 * - Returns one next candle bar
 * - Subtract the latest candle bar
 * - Get all candle bars get so far
 * - Get the read index
 * - Get the available indexes (optional)
 */
pub trait OhlcReplay {
    fn get_state(&self) -> OhlcReplayState;
    fn set_read_index(&mut self, new_read_index: u32) -> Result<u32, EconomyError>;
    fn get_ohlc_up_to_read_index(&mut self, read_index: u32) -> Result<Vec<Ohlc>, EconomyError>;
    
    // fn increase_read_index(&self, interval: u32)
}
