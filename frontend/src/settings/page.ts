import { writable } from "svelte/store";

export enum PageType {
  Account = 0,
  Connections,
  Appearance,
  Bots,
}
export const settings_page = writable(PageType.Account);
