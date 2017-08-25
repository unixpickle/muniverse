(function() {

  var c2 = construct2api;

  var scoreMode = 'distance';
  function rawScore() {
    if (scoreMode === 'distance') {
      return c2.globalVar('Score').data;
    } else {
      return c2.globalVar('Coins').data;
    }
  }

  var gameOver = false;
  var endgameScore = null;
  window.muniverse = {
    init: function(options) {
      scoreMode = options.scoreMode;
      if (['distance', 'coins'].indexOf(scoreMode) < 0) {
        return Promise.reject('invalid scoreMode: ' + scoreMode);
      }
      return init4j().then(function() {
        faketime.pause();

        var rt = cr_getC2Runtime();
        rt.doChangeLayout(rt.layouts.Game);

        faketime.advance(500);

        // Simulate screen tap.
        c2.runActions([
          "877053668669505",
          "690973812839542",
          "963439334977315",
          "301360938312898",
          "2003561706778339",
          "1666160326853771",
          "138928761058736"
        ]);

        // Disable pause button.
        c2.disableActions([
          "4164650708027873",
          "7816534686043383",
          "6050832782901967",
          "3134628972584548",
          "8663479218992549",
          "8029353703779944",
          "762795278645658",
          "1647097791598594",
          "7953746193468714",
          "8978928891585849",
          "7284102521188451",
          "9274071293115874",
          "7232745202374412",
          "4813914032084102",
          "3988927934107422",
          "7316127759972545",
          "8569961392587366",
          "3721949730698611",
          "766582089147191",
          "8257680915782511",
          "2812044842869475",
          "9030942861782774",
          "2848824460608151"
        ]);

        rt.doChangeLayout = () => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = rawScore();
          }
        }
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(gameOver ? endgameScore : rawScore());
    }
  };

})();
