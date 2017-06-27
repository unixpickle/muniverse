(function() {

  var PAUSE_BUTTON_UID = '132';

  // Game depends on a certain framerate.
  faketime.animationFrameRate = 60;

  // Avoid tutorial blurb.
  localStorage.coin = '1';

  var c2 = window.construct2api;

  window.muniverse = {
    init: function(options) {
      return c2.waitStart('game').then(function() {
        faketime.pause();
        var challengeActions = [
          "3405174827741924",
          "2117344253154359",
          "28639723073376",
          "6373392234851472",
          "3802899951408011",
          "6964386149465146",
          "6412880831969221",
          "7556063816109108",
          "3626872888522236",
          "2894652460813217",
          "537275907231722",
          "5017233555978647",
          "3337848642680236",
          "8677996837111956",
          "6866266498403392",
          "1151670455251325"
        ];
        if (options.mode === 'challenge') {
          c2.runActions(challengeActions);
        } else if (options.mode === 'time') {
          c2.runActions(challengeActions.concat([
            "9339683925139618",
            "6189233853686172"
          ]));
        }
        faketime.advance(100);
        cr_getC2Runtime().objectsByUid[PAUSE_BUTTON_UID].visible = false;
        faketime.advance(100);

        // Used to detect game over.
        ['best0', 'best1'].forEach((k) => c2.globalVar(k).setValue(-1));
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      var done = ['best0', 'best1'].some((k) => c2.globalVar(k).getValue() >= 0);
      return Promise.resolve(done);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
