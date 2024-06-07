import { atom } from "jotai";

import {
  Bookmark,
  Message,
  MessageMessageType,
  Thread,
} from "@/__generated__/graphql";

export const threadAtom = atom<Thread | undefined>(undefined);

export const threadMessagesAtom = atom<Message[]>((get) => {
  let messages: Message[] = [];

  get(threadAtom)?.messages?.edges?.forEach((edge) => {
    if (edge?.node) {
      messages.push(edge.node as Message);
    }
  });

  return messages;
});

export const lastThreadMessageAtom = atom<Message | undefined>((get) => {
  let messages = get(threadMessagesAtom);

  if (messages.length > 0) {
    return messages[messages.length - 1];
  }

  return undefined;
});

export const lastThreadMessageTypeAtom = atom<MessageMessageType | undefined>(
  (get) => {
    return get(lastThreadMessageAtom)?.messageType;
  },
);

export const threadBookmarksAtom = atom<Bookmark[]>([]);

export const parentThreadAtom = atom<Thread | undefined>(undefined);

export const parentThreadMessagesAtom = atom<Message[]>((get) => {
  let messages: Message[] = [];

  get(parentThreadAtom)?.messages?.edges?.forEach((edge) => {
    if (edge?.node) {
      messages.push(edge.node as Message);
    }
  });

  return messages;
});
