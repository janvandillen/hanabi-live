// Functions having to do with the "fade" and "critical" features
// (e.g. cards that no longer need to be played are faded,
// cards that become critical get an icon)

import globals from './globals';

// Resets the fade and critical for every card in game
export default function cardStatusChecks() {
  if (globals.animateFast) {
    return;
  }

  for (let i = 0; i <= globals.indexOfLastDrawnCard; i++) {
    const card = globals.deck[i];
    card.setFade();
    card.setCritical();
  }
}
