export function setCachedUserId(email: string, userId: string): void {
  global.cacheUser.set(email, userId);
}

export function getCachedUserId(email: string): string | undefined {
  return global.cacheUser.get(email);
}
