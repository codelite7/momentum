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
  const { children, onScrollDown, onScrollUp, gap, ...rest } = props;
  let flexGap = gap ? gap : 2;

  return (
    <div className="h-full overflow-y-auto">
      <ScrollShadow
        {...rest}
        onScroll={(e) => onInfiniteScroll(e, onScrollDown, onScrollUp)}
      >
        <div className={`flex flex-col gap-${flexGap}`}>{children}</div>
      </ScrollShadow>
    </div>
  );
}
