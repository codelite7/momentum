"use client";

import { useState, useEffect } from "react";
import { Button, Card, CardBody, CardFooter, Tooltip } from "@nextui-org/react";
import { redirect, useRouter } from "next/navigation";
import { useSuspenseQuery } from "@apollo/client";
import { Link } from "@nextui-org/link";

import useAccountSettingsModalStore from "@/stores/account-settings-modal-store";
import useSearchModalStore from "@/stores/search-modal-store";
import ThreadButtons from "@/components/left-sidebar/ThreadButtons";
import { gql } from "@/__generated__";

export const sidebarThreadsQuery = gql(/* GraphQL */ `
  query sidebarThreadsQuery {
    user {
      id
      email
      threads(first: 50, orderBy: { field: LAST_VIEWED_AT, direction: DESC }) {
        totalCount
        pageInfo {
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
        edges {
          node {
            id
            name
          }
        }
      }
    }
  }
`);

export default function LeftSidebar() {
  const router = useRouter();
  const { data, error } = useSuspenseQuery(sidebarThreadsQuery);

  if (error) {
    console.log("left-sidebar error", error);
  } else {
  }
  const toggleAccountSettings = useAccountSettingsModalStore(
    (state: any) => state.toggle,
  );
  const toggleSearch = useSearchModalStore((state: any) => state.toggle);
  const [collapsed, setCollapsed] = useState(false);
  const [width, setWidth] = useState("w-52");
  const user = data?.user;

  useEffect(() => {
    let width = "w-52";

    if (collapsed) {
      width = "w-16";
    }
    setWidth(width);
  }, [collapsed]);

  return (
    <>
      <Card
        className={`h-full ${width} rounded-none !transition-all !duration-1000`}
      >
        <CardBody className="h-full overflow-hidden p-0">
          {collapsed ? (
            <>
              <div className="h-full p-2">
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
                    <img alt="stratifi" className="h-4 w-4" src="/logo.png" />
                  </Button>
                </Tooltip>

                <Tooltip showArrow content="Search" placement="right">
                  <Button
                    isIconOnly
                    className="w-full mt-2"
                    color="primary"
                    size="sm"
                    variant="bordered"
                    onPress={toggleSearch}
                  >
                    <img alt="search" className="h-4 w-4" src="/search.svg" />
                  </Button>
                </Tooltip>
                <Tooltip showArrow content="New thread" placement="right">
                  <Button
                    isIconOnly
                    as={Link}
                    className="w-full mt-2"
                    color="primary"
                    href="/thread/new"
                    size="sm"
                    onPress={() => {
                      redirect("/thread/new");
                    }}
                  >
                    <img
                      alt="new thread"
                      className="h-4 w-4"
                      src="/layers.svg"
                    />
                  </Button>
                </Tooltip>
              </div>
            </>
          ) : (
            <>
              <div className="p-2">
                <div className="w-full flex flex-row justify-between items-center mb-2">
                  <div className="flex items-center">
                    <img
                      alt="stratifi logo"
                      className="mr-2"
                      height="20"
                      src="/logo.png"
                      width="20"
                    />
                    <span className="tracking-widest">STRATIFI</span>
                  </div>
                  <Tooltip showArrow content="Collapse" placement="right">
                    <Button
                      isIconOnly
                      className="min-w-6 w-6 h-6 rounded"
                      variant="light"
                      onPress={() => {
                        setCollapsed(true);
                        setWidth("w-16");
                      }}
                    >
                      <i className="pi pi-chevron-left" />
                    </Button>
                  </Tooltip>
                </div>
                {/* search box */}
                <Button
                  fullWidth
                  className="flex flex-row start mb-2"
                  color="primary"
                  size="sm"
                  variant="bordered"
                  onPress={toggleSearch}
                >
                  <div className="w-full flex start text-white rounded items-center">
                    <img alt="search" className="mr-2" src="/search.svg" />
                    Search
                  </div>
                </Button>
                {/* new thread button*/}
                <Button
                  fullWidth
                  as={Link}
                  color="primary"
                  href="/thread/new"
                  size="sm"
                  startContent={<img alt="new thread" src="/layers.svg" />}
                >
                  New thread
                </Button>
              </div>
              {/* thread buttons */}
              <ThreadButtons threads={user.threads} />
            </>
          )}
        </CardBody>
        <CardFooter className="bg-black rounded-none">
          {collapsed ? (
            <>
              <Tooltip showArrow content="Account settings" placement="right">
                <Button
                  isIconOnly
                  className="w-full"
                  // className="min-w-8 w-8 h-8 rounded mr-4"
                  color="primary"
                  size="sm"
                  onPress={toggleAccountSettings}
                >
                  <img
                    alt="avatar"
                    className="h-4 w-4"
                    src="/default-user-avatar.svg"
                  />
                </Button>
              </Tooltip>
            </>
          ) : (
            <Button
              isIconOnly
              className="min-w-8 w-8 h-8 rounded mr-4"
              color="primary"
              size="sm"
              onPress={toggleAccountSettings}
            >
              <img
                alt="avatar"
                className="h-4 w-4"
                src="/default-user-avatar.svg"
              />
            </Button>
          )}
        </CardFooter>
      </Card>
    </>
  );
}
