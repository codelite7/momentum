"use client";

import { Button, Card, CardBody, Tooltip } from "@nextui-org/react";
import { useEffect, useState } from "react";

import Bookmarks from "@/components/right-sidebar/bookmarks";

export function RightSidebar() {
  const [collapsed, setCollapsed] = useState(false);
  const [width, setWidth] = useState("w-64");

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
      className={`h-full ${width} rounded-none !transition-all !duration-1000`}
    >
      <CardBody className="p-2">
        {collapsed ? (
          <>
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
          </>
        ) : (
          <>
            {/* header */}
            <div className="flex items-center mb-4">
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
            {/* bookmarks search*/}
            <Button
              fullWidth
              className="flex flex-row start mb-2 mt-4"
              color="primary"
              size="sm"
              variant="bordered"
            >
              <div className="w-full flex start text-white rounded items-center">
                <img alt="search" className="mr-2" src="/search.svg" />
                Search all bookmarks
              </div>
            </Button>
            {/* thread bookmarks */}
            <Bookmarks
              params={{ id: "1ff4ad0d-4355-4123-a57a-9c76cac135e0" }}
            />
          </>
        )}
      </CardBody>
    </Card>
  );
}
