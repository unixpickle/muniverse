(function() {

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('score2').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        construct2api.runActions([
          "4521313383244316", "7485039192927952", "6835872615093203", "5871818170303755",
          "581177690993926", "7050022383246375", "485031209287812", "2133515256146083",
          "5864187521024456", "7869777241359649", "7074937924258305", "884893382078333",
          "5380214772043478", "1555113621046625", "5527837649467908", "8828277125267639",
          "2872872813096689", "7448690018834042", "8138971325149974", "5447975039460945",
          "1069160109928945", "2056246146470745", "6323235292249072", "4294073400888452",
          "8665408678789103", "7419199717070001", "7810055040365703", "8742566763900119",
          "9140018789684040", "9497414451094532", "5265569669541993"
        ]);
        faketime.advance(1000);
        cr_getC2Runtime().doChangeLayout = () => {
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
