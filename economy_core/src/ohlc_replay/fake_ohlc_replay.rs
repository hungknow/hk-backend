use std::ops::Sub;

use crate::{EconomyError, Ohlc, OhlcReplay, OhlcReplayState, Next};

pub struct FakeOhlcReplay {
    state: OhlcReplayState, // Maximum state
    vec: Vec<Ohlc>,

    // nextValue: &'r mut dyn Next<Ohlc>,
}

impl FakeOhlcReplay {
    pub fn new_random(target_available_count: usize) -> FakeOhlcReplay {
        let mut vec = Vec::new();
        for n in 1..target_available_count + 1 {
            vec.push(Ohlc {
                close_time: chrono::Utc::now().sub(chrono::Duration::seconds(
                    (target_available_count - n) as i64,
                )),
                high: 0.0,
                open: 0.0,
                close: 0.0,
                low: 0.0,
                volume: Some(0.0),
                trade_count: Some(0.0),
            });
        }

        FakeOhlcReplay {
            state: OhlcReplayState {
                available_count: vec.len(),
                data_exhausted: true,
            },
            vec: vec,
        }
    }

    // pub fn new_from_csv_parser(next_value: impl Next<Ohlc>) -> FakeOhlcReplay {
    //     FakeOhlcReplay {
    //         state: OhlcReplayState {
    //             available_count: 0,
    //             data_exhausted: false,
    //         },
    //         vec: Vec::new(),
    //     }
    // }
}

impl OhlcReplay for FakeOhlcReplay {
    fn get_state(&self) -> OhlcReplayState {
        self.state
    }

    fn get_ohlc_up_to_read_index(&mut self, count: usize) -> Result<Vec<Ohlc>, EconomyError> {
        if count > self.state.available_count {
            return Err(EconomyError::OutOfRange);
        }
        let data = self.vec[0..count as usize].to_vec();
        self.state.available_count = self.vec.len();
        Ok(data)
    }

    fn get_ohlc_at_index(&mut self, read_index: usize) -> Result<Ohlc, EconomyError> {
        if read_index >= self.state.available_count {
            return Err(EconomyError::OutOfRange);
        }

        match self.vec.get(read_index as usize) {
            Some(ohlc) => Ok(ohlc.clone()),
            None => Err(EconomyError::InvalidParameter),
        }
    }
}

#[cfg(test)]
mod tests {
    use core::panic;

    use super::*;

    #[test]
    fn ohlc_replay_new_random() {
        let max_available: usize = 2;
        let mut fake_ohlc_replay = FakeOhlcReplay::new_random(max_available);
        assert_eq!(fake_ohlc_replay.vec.len(), max_available as usize);

        let mut state = fake_ohlc_replay.get_state();
        assert_eq!(state.available_count, max_available);
        assert_eq!(state.data_exhausted, true);

        // It's OK to read data at index 0
        let ohlc_at_index0: Result<Ohlc, EconomyError> = fake_ohlc_replay.get_ohlc_at_index(0);
        match ohlc_at_index0 {
            Ok(_) => {}
            Err(err) => {
                panic!("{}", err)
            }
        }
        // After reading the value, the length of vector is still the same
        assert_eq!(fake_ohlc_replay.vec.len(), max_available as usize);

        // It's OK to read data at index 1
        let mut ohlc_at_index1 = fake_ohlc_replay.get_ohlc_at_index(1);
        match ohlc_at_index1 {
            Ok(_) => {}
            Err(err) => {
                panic!("{}", err)
            }
        }

        // assert_eq!(fake_ohlc_replay.set_read_index(1), Ok(0));
        let mut ohlc: Vec<Ohlc> = fake_ohlc_replay.get_ohlc_up_to_read_index(1).unwrap();
        assert_eq!(ohlc.len(), 1);

        // It's still OK to read data at index 1 after setting read_index
        ohlc_at_index1 = fake_ohlc_replay.get_ohlc_at_index(1);
        match ohlc_at_index1 {
            Ok(_) => {}
            Err(err) => {
                panic!("{}", err)
            }
        }

        ohlc = fake_ohlc_replay.get_ohlc_up_to_read_index(2).unwrap();
        assert_eq!(ohlc.len(), 2);

        // After reading the HOLC up to the read index,
        // the value of read _index ist still the same
        state = fake_ohlc_replay.get_state();

        ohlc = fake_ohlc_replay
            .get_ohlc_up_to_read_index(max_available)
            .unwrap();
        assert_eq!(ohlc.len(), max_available as usize);
        state = fake_ohlc_replay.get_state();

        // Set the read index to a specific index
        state = fake_ohlc_replay.get_state();
        assert_eq!(state.available_count, max_available);
        assert_eq!(ohlc.len(), max_available as usize);
    }
}
