use serde::{Deserialize, Serialize};

/// Communicates the state of the [`Feed`] as well as the next event.
#[derive(Clone, Eq, PartialEq, PartialOrd, Debug, Deserialize, Serialize)]
pub enum Feed<Event> {
    Next(Event),
    Unhealthy,
    Finished,
}
