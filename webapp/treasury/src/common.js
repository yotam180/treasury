export function processDate(dateStr) {
  const date = new Date(dateStr);
  return `${date.getDay()}/${date.getMonth()}/${date.getFullYear()}`;
}
