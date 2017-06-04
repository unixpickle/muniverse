(function() {

  var SPLASH_TIMEOUT = 60000;

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('Game').then(function() {
        return pollAndWait(SPLASH_TIMEOUT, function() {
          var layout = cr_getC2Runtime().running_layout;
          return layout.layers[1].instances.length == 11;
        });
      }).then(function() {
        var startActions = [
          "4774391839033219",
          "4911066361684777",
          "8998198662723024",
          "9019906187189800",
          "740391283959218",
          "6979027526638077",
          "2829065343360993",
          "1979720213788508",
          "1989928345543413",
          "6105618985824789"
        ];
        c2.runActions(startActions);

        window.faketime.pause();
        window.faketime.advance(40);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(c2.globalVar('speed').getValue() === 0);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
