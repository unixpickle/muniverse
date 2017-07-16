(function() {

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.Game);
        faketime.advance(100);

        ['mousedown', 'mouseup'].forEach((evt) => {
          var e = new jQuery.Event(evt);
          e.pageX = $('#gameCanvas').offset().left + $('#gameCanvas').width()/2;
          e.pageY = $('#gameCanvas').offset().left + $('#gameCanvas').height()*2/3;
          jQuery('#gameCanvas').trigger(e);
        });

        // Detect game over.
        c2.globalVar('BestScore').data = -1;
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(c2.globalVar('BestScore').data >= 0);
    },
    score: function() {
      const best = c2.globalVar('BestScore').data;
      return Promise.resolve(best >= 0 && best || c2.globalVar('Score').data);
    }
  };

})();
