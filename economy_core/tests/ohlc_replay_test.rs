use csv::ReaderBuilder;
use economy_core::{FakeOhlcReplay, OhlcReplay, RelativeStrengthIndex, Next, OhlcCsvParser};

#[test]
fn ohlc_replay_features() {
    let target_available_count = 2;
    // Connect Ohlc Replay to the Ohlc Data
    let mut fake_ohlc_replay: FakeOhlcReplay = FakeOhlcReplay::new_random(target_available_count);

    let data = "\
    city;country;pop
    Boston;United States;4628910
    ";
    let mut rdr: csv::Reader<&[u8]> = ReaderBuilder::new()
    .delimiter(b';')
    .from_reader(data.as_bytes());
    let csvParser = OhlcCsvParser::parse()

    let mut csv_fake_ohlc_replay: FakeOhlcReplay = FakeOhlcReplay::new_random(target_available_count);

    let mut fake_ohlc_replay_state = fake_ohlc_replay.get_state();
    // println!("{:?}", fake_ohlc_replay_state);
    assert_eq!(fake_ohlc_replay_state.available_count, target_available_count);

    // Indicator receive the new Ohlc data from "Ohlc Replay", process it,
    let mut rsiIndicator = RelativeStrengthIndex::new(14).unwrap();

    let mut current_read_index = 0;

    loop {
        fake_ohlc_replay_state = fake_ohlc_replay.get_state();
        // println!("{:?}", fake_ohlc_replay_state);
        if current_read_index >= fake_ohlc_replay_state.available_count {
            break;
        }

        let ohlc = match fake_ohlc_replay.get_ohlc_at_index(current_read_index) {
            Ok(ohlc) => ohlc,
            Err(err) => panic!("{}", err),
        };
        current_read_index += 1;
        let rsiValue = rsiIndicator.next(ohlc.high);
    }

}
