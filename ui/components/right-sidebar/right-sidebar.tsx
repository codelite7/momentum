"use client";

import { Button, Card, CardBody, Tooltip } from "@nextui-org/react";
import { useEffect, useState } from "react";
import { useAtomValue } from "jotai/index";

import Bookmarks from "@/components/right-sidebar/bookmarks";
import { threadAtom } from "@/state/atoms";

export function RightSidebar() {
  const [collapsed, setCollapsed] = useState(false);
  const [width, setWidth] = useState("w-64");
  const thread = useAtomValue(threadAtom);

  useEffect(() => {
    let width = "w-64";

    if (collapsed) {
      width = "w-16";
    }
    setWidth(width);
  }, [collapsed]);

  return (
    <Card
      disableAnimation
      className={`h-full ${width} rounded-none !transition-all !duration-1000 bg-background border-l-1 border-default`}
    >
      <CardBody className="p-0">
        {collapsed ? (
          <div className="px-4 pt-2">
            {/* expand button*/}
            <Tooltip showArrow content="Expand" placement="right">
              <Button
                isIconOnly
                className="w-full"
                color="primary"
                size="sm"
                variant="light"
                onPress={() => {
                  setCollapsed(!collapsed);
                }}
              >
                <i className="pi pi-bookmark" />
              </Button>
            </Tooltip>
            {/* bookmark search button*/}
            <Tooltip showArrow content="Search" placement="right">
              <Button
                isIconOnly
                className="w-full mt-2"
                color="primary"
                size="sm"
                variant="bordered"
                onPress={() => {}}
              >
                <img alt="search" className="h-4 w-4" src="/search.svg" />
              </Button>
            </Tooltip>
          </div>
        ) : (
          <>
            {/* header */}
            <div className="flex items-center min-h-11 w-full bg-card-background border-b-1 border-default px-4">
              <Tooltip showArrow content="Collapse" placement="left">
                <Button
                  isIconOnly
                  className="min-w-6 w-6 h-6 rounded mr-2"
                  variant="light"
                  onPress={() => {
                    setCollapsed(true);
                    setWidth("w-16");
                  }}
                >
                  <i className="pi pi-chevron-right" />
                </Button>
              </Tooltip>
              Bookmarks
            </div>
            <div className="px-4 mt-2">
              {/* bookmarks search*/}
              {/*<Button*/}
              {/*  fullWidth*/}
              {/*  className="flex flex-row start my-2"*/}
              {/*  color="primary"*/}
              {/*  size="sm"*/}
              {/*  variant="bordered"*/}
              {/*>*/}
              {/*  <div className="w-full flex start text-white rounded items-center">*/}
              {/*    <img alt="search" className="mr-2" src="/search.svg" />*/}
              {/*    Search all bookmarks*/}
              {/*  </div>*/}
              {/*</Button>*/}
              {thread && <Bookmarks />}
            </div>
          </>
        )}
      </CardBody>
    </Card>
  );
}
