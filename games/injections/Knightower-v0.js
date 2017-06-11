(function() {

  var c2 = window.construct2api;

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return c2.waitStart('game').then(function() {
        cr_getC2Runtime().actsBySid["9591441365881388"].run = () => gameOver = true;

        window.faketime.pause();
        window.construct2api.runActions([
          "9001497584259579",
          "9941433678582528",
          "7522014314257657",
          "4745377136575544",
          "8579199023683658",
          "7012221537726872",
          "4649151013729039",
          "9689694779777128",
          "4604904302237823",
          "1023914123585311",
          "630463858130439",
          "6262138772026519",
          "1946929135193365",
          "6379886936667443",
          "2558672227548044",
          "6460162894620466",
          "4804170843859204"
        ]);
        window.faketime.advance(1000);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
