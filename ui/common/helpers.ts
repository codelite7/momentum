export function getSentAtDisplay(createdAt: string): string {
  let sentAtDisplay = "";

  if (createdAt) {
    const sentAtDate = new Date(createdAt);
    const sentToday =
      new Date(createdAt).setHours(0, 0, 0, 0) ==
      new Date().setHours(0, 0, 0, 0);

    sentAtDisplay = sentToday
      ? sentAtDate.toLocaleTimeString()
      : sentAtDate.toLocaleString();
  }

  return sentAtDisplay;
}
