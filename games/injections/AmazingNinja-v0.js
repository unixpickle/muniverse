(function() {

  var c2 = construct2api;

  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        ['mousedown', 'mouseup'].forEach((evt) => {
          var e = new jQuery.Event(evt);
          e.pageX = 10;
          e.pageY = 10;
          jQuery(document).trigger(e);
        });
        faketime.advance(100);

        // Used to detect game over.
        c2.globalVar('BestScore').data = -1;
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(c2.globalVar('BestScore').data >= 0);
    },
    score: function() {
      var bs = c2.globalVar('BestScore').data;
      return Promise.resolve(bs >= 0 ? bs : c2.globalVar('Score').data);
    }
  };

})();
