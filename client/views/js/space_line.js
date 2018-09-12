function loopSpace(count) {
  var loopSpaceWrite = "";
  for (i = 0; i <= count; i++) {
    loopSpaceWrite += "&nbsp;";
  }

  return loopSpaceWrite;
}

calculateSpace = length => {
  space = 0;
  switch (length) {
    case 1:
    case 2:
      space = 53 - length;
      break;
    case 3:
    case 4:
      space = 50 - length;
      break;
    case 5:
      space = 48 - length;
      break;
    case 6:
      space = 47 - length;
      break;
    case 7:
    case 8:
      space = 46 - length;
      break;
    case 9:
      space = 44 - length;
      break;
    case 10:
      space = 43 - length;
      break;
    case 11:
      space = 41 - length;
      break;
    case 12:
    case 13:
      space = 40 - length;
      break;
    case 14:
    case 15:
      space = 39 - length;
      break;
    case 16:
    case 17:
      space = 38 - length;
      break;
    default:
      break;
  }
  return space;
};
