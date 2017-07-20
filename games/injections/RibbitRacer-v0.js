(function() {

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('score').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();

        // Skip gradual transition in.
        faketime.advance(2000);

        // Hit start button.
        c2.runActions([
          "3345439722153657", "634593180451865", "5687677633976457",
          "5709964739400251", "180636297805455", "1066143644264542",
          "5959952595053187", "6682179108170979", "2932359162566739",
          "9240090261002136", "3205047195370838", "510402835649117",
          "6391216278270219", "8153321935907514", "1951743421271061",
          "7289499067457064", "173798993242823", "650862276778631"
        ]);
        faketime.advance(100);

        // Tap to start.
        c2.runActions([
          "4895184116071378", "8770543717606078", "9850427451177112",
          "35831634136616", "763682059559291", "859359858457399"
        ]);
        faketime.advance(100);

        c2.globalVar('gameover').setValue = () => {
          gameOver = true;
          endgameScore = rawScore();
        };
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
