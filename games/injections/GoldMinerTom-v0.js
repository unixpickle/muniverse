(function() {

  var c2 = window.construct2api;

  // Prevent state from propagating between episodes.
  window.indexedDB.deleteDatabase('localforage');

  window.muniverse = {
    init: function() {
      return c2.waitStart('menu').then(function() {
        window.faketime.pause();

        var startActions = [
          "221934939143096",
          "616907927503173",
          "281104572458521",
          "862159693914130",
          "783999930970336",
          "562417132443266",
          "640271430658423"
        ]
        c2.runActions(startActions);

        window.faketime.advance(100);

        var playActions = [
          "420824669656968",
          "445565207731627",
          "101637426543737",
          "666227923162335",
          "652852076898327",
          "497148326540854",
          "952846340368005",
          "923302072648596",
          "562198753540527",
          "410402783117193",
          "183694449540330",
          "517752215114989",
          "427288469060264",
          "334528067134172",
          "900738740267264",
          "790565667788517",
          "384809854886198",
          "313104697295777",
          "170665595893705"
        ];
        c2.runActions(playActions);

        var pauseActions = [
          "688965941005344",
          "401874433656547",
          "192424283039903",
          "645830181526677",
          "142680415588832",
          "433138071525098",
          "365119968703232",
          "132100401413251",
          "915771881554215",
          "997402763455660",
          "743562109393775",
          "730089550038174",
          "543132899935963",
          "873199439603100",
          "559903953657638",
          "457654427117272",
          "445932609793024",
          "661842479903847",
          "858248818778987"
        ];
        c2.disableActions(pauseActions);

        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(c2.globalVar('popup').getValue() === 1 ||
        c2.globalVar('pause').getValue() === 1);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
