var GetQS = function(param) {
  var sPageURL = window.location.search.substring(1),
    sURLVariables = sPageURL.split(/[&||?]/),
    res;
  for (var i = 0; i < sURLVariables.length; i += 1) {
    var paramName = sURLVariables[i],
      sParameterName = (paramName || '').split('=');

    if (sParameterName[0] === param) {
      res = sParameterName[1];
    }
  }
  return res;
}
var SetQS = function(param, paramVal) {
  var newAdditionalURL = "";
  var tempArray = window.location.toString().split("?");
  var additionalURL = tempArray[1];
  var temp = "";
  if (additionalURL) {
    tempArray = additionalURL.split("&");
    console.log(tempArray, tempArray.length);
    for (var i = 0; i < tempArray.length; i++) {
      console.log(tempArray[i].split('=')[0]);
      if (tempArray[i].split('=')[0] != param) {
        console.log(tempArray[i].split('=')[0], tempArray[i].split('=')[1]);
        newAdditionalURL += temp + tempArray[i];
        temp = "&";
        console.log(temp);
      }
    }
  }

  var rows_txt = temp + "" + param + "=" + paramVal;
  console.log("?" + newAdditionalURL + rows_txt);
  window.history.replaceState('', '', "?" + newAdditionalURL + rows_txt);
}

function hide(id) {
  let login = document.getElementById(id);
  login.style.display = 'none';
}

function show(id) {
  let login = document.getElementById(id);
  login.style.display = 'block';
}

function jump(h) {
  let a = document.getElementById(h)
  if (a != undefined) {
    let top = a.offsetTop; //Getting Y of target element
    window.scrollTo(0, top); //Go there directly or some transition
  }
}

window.onload = function() {}
