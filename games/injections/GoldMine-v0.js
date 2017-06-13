(function() {

  var c2 = window.construct2api;

  // Prevent state from propagating between episodes.
  window.indexedDB.deleteDatabase('localforage');

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return c2.waitStart('Menu').then(function() {
        window.faketime.pause();

        var startActions = [
          "6291856493950302",
          "7213866823074631",
          "5272823363511825",
          "6799124326211135",
          "6824313990361322",
          "2203714347427264",
          "470875311115557"
        ];
        c2.runActions(startActions);

        window.faketime.advance(100);

        var pauseActions = [
          "6734604015435425",
          "7341606590179908",
          "6377480576975344",
          "9901628227923096",
          "1244203631251041",
          "5447824283000052",
          "691809414589615",
          "385607789018749",
          "2559597143450936",
          "8517287343113593",
          "6132034817592136",
          "5117947770664292",
          "5961759676768692",
          "5543496868371069",
          "7728096130810856",
          "8564180171143871",
          "4834061956530859",
          "2923747017529466"
        ];
        c2.disableActions(pauseActions);

        var gameOverActs = [
          '732797195051172',
          '6337302889982321',
          '5117947770664292'
        ];
        gameOverActs.forEach((x) => c2.hijackAction(x, () => gameOver = true));

        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score2').getValue());
    }
  };

})();
