import { memo } from "react";
import { useSuspenseQuery } from "@apollo/client";
import { ScrollShadow } from "@nextui-org/react";

import { searchResultsQuery } from "@/graphql-queries/queries";
import MessageResult from "@/components/search/message-result";
import { Message } from "@/__generated__/graphql";

interface props {
  query: string;
}

export default memo(function SearchResults({ query }: props) {
  const { data } = useSuspenseQuery(searchResultsQuery, {
    variables: {
      query: query,
    },
  });

  return (
    <>
      {data && data.messages.totalCount > 0 ? (
        <>
          <span>Matches: {data.messages.totalCount}</span>
          <ScrollShadow hideScrollBar className="max-h-96">
            {data.messages.edges?.map((edge) => {
              return (
                <>
                  {edge?.node && (
                    <MessageResult
                      key={edge?.node?.id}
                      message={edge?.node as Message}
                      query={query}
                    />
                  )}
                </>
              );
            })}
          </ScrollShadow>
        </>
      ) : (
        <div className="w-full text-center">No results</div>
      )}
    </>
  );
});
