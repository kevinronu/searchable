// const options: Intl.DateTimeFormatOptions = {
//   weekday: "long",
//   year: "numeric",
//   month: "long",
//   day: "numeric",
//   timeZone: "UTC",
//   hour: "2-digit",
//   minute: "2-digit",
// };

export function parseDate(
  date: Date | string,
  options: Intl.DateTimeFormatOptions
): string {
  if (typeof date === "string") {
    date = new Date(date);
  }

  if (date instanceof Date && !isNaN(date.getTime())) {
    return date.toLocaleDateString("en-US", options);
  } else {
    return "Invalid Date";
  }
}
