use crate::{Ohlc, EconomyError, Feed};

pub trait Reset {
    fn reset(&mut self);
}

/// Return the period used by the indicator.
pub trait Period {
    fn period(&self) -> usize;
}

pub trait Next<T> {
    type Output;
    fn next(&mut self, input: T) -> Self::Output;
}

/// Generates the next `Event`. Acts as the system heartbeat.
pub trait MarketGenerator<Event> {
    /// Return the next market `Event`.
    fn next(&mut self) -> Feed<Event>;
}

/// Close price of a particular period.
pub trait Close {
   fn close(&self) -> f64; 
}

// pub trait OhlcCsvParser {
//     fn parse(reader: csv::Reader<&[u8]>) -> Option<Result<Vec<Ohlc>, EconomyError>>;
// }

pub trait OhlcReader {
    fn get(&self, symbol: &str, query: &str) -> Vec<Ohlc>;
}

#[derive(Clone, Copy, Debug)]
pub struct OhlcReplayState {
    pub available_count: usize,
    pub data_exhausted: bool
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
    fn get_ohlc_up_to_read_index(&mut self, read_index: usize) -> Result<Vec<Ohlc>, EconomyError>;
    fn get_ohlc_at_index(&mut self, read_index: usize) -> Result<Ohlc, EconomyError>;
}
