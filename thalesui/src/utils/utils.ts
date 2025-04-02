
export const roundToTwoDecimals = (val: number): number => {
  return Math.round(val * 100) / 100;
};

export function formatNumberWithCommas(value?: number) {
  return roundToTwoDecimals(value || 0) // round to 2 decimal places first.
    .toFixed(2) // make sure the number is always showing 2dp. i.e. 1 => 1.00, 1.1 => 1.10
    .replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}
