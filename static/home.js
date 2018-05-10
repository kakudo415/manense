'use strict';
var balanceHTML = document.getElementById('balance');
var editHTML = document.getElementById('edit');
var editNameHTML = document.getElementById('editName');
var editIncomeHTML = document.getElementById('editIncome');
var editTimeHTML = document.getElementById('editTime');

var currentID;
function exMenu(exID, on) {
  currentID = exID;
  if (on) {
    if (currentID.length > 0) {
      var AJAX = new XMLHttpRequest();
      AJAX.open('POST', '/info', true);
      AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
      AJAX.send('expense-uuid=' + exID);
      AJAX.onreadystatechange = () => {
        if (AJAX.readyState == 4 && AJAX.status == 200) {
          var data = JSON.parse(AJAX.responseText);
          editNameHTML.value = data.Name;
          editIncomeHTML.value = data.Income;
          editTimeHTML.value = data.Time.substr(0, 10);
          editHTML.style.display = 'block';
        }
      }
    } else {
      editTimeHTML.value = nowString();
      editHTML.style.display = 'block';
    }
  } else {
    editHTML.style.display = 'none';
  }
}

function exSave() {
  var AJAX = new XMLHttpRequest();
  if (currentID.length > 0) {
    AJAX.open('POST', '/update', true);
    AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    AJAX.send('expense-uuid=' + currentID + '&expense-name=' + editNameHTML.value + '&expense-income=' + editIncomeHTML.value + '&expense-time=' + editTimeHTML.value);
    AJAX.onreadystatechange = () => {
      if (AJAX.readyState == 4 && AJAX.status == 200) {
        if (editIncomeHTML.value >= 0) {
          document.getElementById(currentID).innerHTML = `<span>${editNameHTML.value}</span><span>${editTimeHTML.value}</span><span>${editIncomeHTML.value} 円</span>`;
        } else {
          document.getElementById(currentID).innerHTML = `<span>${editNameHTML.value}</span><span>${editTimeHTML.value}</span><span class="minus">${editIncomeHTML.value} 円</span>`;
        }
        if (Number(AJAX.responseText) >= 0) {
          balanceHTML.innerHTML = `残高 : ${AJAX.responseText} 円`;
        } else {
          balanceHTML.innerHTML = `残高 : <span class="minus">${AJAX.responseText} 円</span>`;
        }
        exMenu('', false);
      }
    }
  } else {
    AJAX.open('POST', '/new', true);
    AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    AJAX.send('expense-name=' + editNameHTML.value + '&expense-income=' + editIncomeHTML.value + '&expense-time=' + editTimeHTML.value);
    AJAX.onreadystatechange = () => {
      if (AJAX.readyState == 4 && AJAX.status == 200) {
        var data = JSON.parse(AJAX.responseText);
        if (editIncomeHTML.value >= 0) {
          document.getElementById('exs').innerHTML = `<a id=${data.uuid} class="ex" onclick="exMenu('${data.uuid}', true);"><span>${editNameHTML.value}</span><span>${editTimeHTML.value}</span><span>${editIncomeHTML.value} 円</span></a>` + document.getElementById('exs').innerHTML;
        } else {
          document.getElementById('exs').innerHTML = `<a id=${data.uuid} class="ex" onclick="exMenu('${data.uuid}', true);"><span>${editNameHTML.value}</span><span>${editTimeHTML.value}</span><span class="minus">${editIncomeHTML.value} 円</span></a>` + document.getElementById('exs').innerHTML;
        }
        if (Number(data.balance) >= 0) {
          balanceHTML.innerHTML = `残高 : ${data.balance} 円`;
        } else {
          balanceHTML.innerHTML = `残高 : <span class="minus">${data.balance} 円</span>`;
        }
        exMenu('', false);
      }
    }
  }
}

function exErase() {
  var AJAX = new XMLHttpRequest();
  AJAX.open('POST', '/erase', true);
  AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
  AJAX.send('expense-uuid=' + currentID);
  AJAX.onreadystatechange = () => {
    if (AJAX.readyState == 4 && AJAX.status == 200) {
      document.getElementById(currentID).style.display = 'none';
      if (Number(AJAX.responseText) >= 0) {
        balanceHTML.innerText = '残高 : ' + AJAX.responseText + ' 円';
      } else {
        balanceHTML.innerHTML = `残高 : <span class="minus">${AJAX.responseText} 円</span>`;
      }
      exMenu('', false);
    }
  }
}

function nowString() {
  var date = new Date();
  var year = date.getFullYear();
  var month = date.getMonth() + 1;
  var day = date.getDate();
  function digit(n, d) {
    n += '';
    if (n.length < d) {
      n = '0' + n;
    }
    return n;
  }
  return digit(year, 4) + '-' + digit(month, 2) + '-' + digit(day, 2);
}
