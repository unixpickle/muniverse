(function() {

  var ENDGAME_TEXTURE = 'images/scoreboard-sheet0.png';

  function rawScore() {
    return construct2api.globalVar('score').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        // Hit play button.
        construct2api.runActions([
          "435483355207331", "349227045446551", "992860352528434",
          "689145158502307", "953210435757187", "299184955372865",
          "363307062568762", "799190037004330", "194060064187376"
        ]);
        faketime.advance(100);

        // Similar endgame logic as NinjaRun-v0.
        var objsByUid = cr_getC2Runtime().objectsByUid;
        var objects = Object.keys(objsByUid).map((x) => objsByUid[x]);
        var views = objects.filter((x) => x.curFrame);
        var dialog = views.find((x) => x.curFrame.texture_file === ENDGAME_TEXTURE);
        dialog.set_bbox_changed = () => {
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
  }

})();
