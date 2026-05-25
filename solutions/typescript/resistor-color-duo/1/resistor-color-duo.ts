const colorCodes: Record<string, number> = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9,
};

export function decodedValue(colors: string[]): number {
  return Number(
    colors
      .slice(0, 2)
      .map(color => colorCodes[color])
      .join("")
  );
}