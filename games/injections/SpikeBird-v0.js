(function() {

  var c2 = construct2api;

  // Game difficulty depends on framerate.
  faketime.animationFrameRate = 60;

  function rawScore() {
    return c2.globalVar('Score').getValue();
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        var rt = cr_getC2Runtime();
        rt.doChangeLayout(rt.layouts.Level1);
        faketime.advance(500);

        rt.doChangeLayout = () => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = rawScore();
          }
        };

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
          "373125842787723",
          "8257680915782511",
          "2812044842869475",
          "9030942861782774",
          "2848824460608151",
          "7235619827719895"
        ]);
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
