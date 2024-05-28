import { ScrollShadow } from "@nextui-org/react";

export function onInfiniteScroll(
  e: any,
  onScrollDown: () => {},
  onScrollUp: () => {},
) {
  const { scrollTop, scrollHeight, clientHeight } = e.target;

  if (scrollTop == 0 && onScrollUp) {
    onScrollUp();
  } else if (scrollTop + clientHeight == scrollHeight && onScrollDown) {
    onScrollDown();
  }
}

export default function InfiniteScroller(props: any) {
  const { children, onScrollDown, onScrollUp, ...rest } = props;

  return (
    <ScrollShadow
      {...rest}
      onScroll={(e) => onInfiniteScroll(e, onScrollDown, onScrollUp)}
    >
      {children}
    </ScrollShadow>
  );
}
