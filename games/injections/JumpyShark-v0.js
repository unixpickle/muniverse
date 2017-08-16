(function() {

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('Score').getValue();
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function(options) {
      var diffToLayout = {
        easy: 'LEVEL_1',
        normal: 'LEVEL2',
        hard: 'LEVEL3'
      };
      var layout = diffToLayout[options.difficulty];
      return init4j().then(function() {
        faketime.pause();
        var rt = cr_getC2Runtime();
        rt.doChangeLayout(rt.layouts[layout]);
        faketime.advance(100);

        construct2api.globalVar('totalDies').setValue = () => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = rawScore();
          }
        };

        // Disable pause button.
        c2.disableActions([
          // Actions for "Easy"
          "6324419816447827",
          "9440042770210724",
          "2103676503851129",
          "4978914420341208",
          "230000791487875",
          "3933455471356607",
          "6365624868362285",
          "789895677327953",
          "637264527486808",

          // Actions for "Normal"
          "3350063230978441",
          "7433819329526261",
          "4042635628370744",
          "6369096977529827",
          "435701678155945",
          "3906907880925484",
          "152098321982275",
          "2665371233334221",
          "5072195083894622",

          // Actions for "Hard"
          "6317132712053317",
          "6649200479537104",
          "6963076627141063",
          "7717843659209033",
          "774927667228092",
          "7498740437861102",
          "3018625012749221",
          "6084504435228892",
          "4004665495299978"
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
