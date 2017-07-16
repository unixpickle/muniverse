(function() {

  // Game speed depends on animation frame rate.
  faketime.animationFrameRate = 60;

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.Game);
        faketime.advance(100);

        // Detect game over.
        c2.globalVar('HighScore').data = -1;

        // Disable exit game button.
        c2.disableActions([
          "7020333949850896",
          "7439249625248278",
          "2135357787493767",
          "8480831502092825"
        ]);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(c2.globalVar('HighScore').data >= 0);
    },
    score: function() {
      const best = c2.globalVar('HighScore').data;
      return Promise.resolve(best >= 0 && best || c2.globalVar('Score').data);
    }
  };

})();
