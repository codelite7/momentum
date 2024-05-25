import {Button} from "@nextui-org/button";
import {useState} from "react";
import ThreadButton from "@/components/left-sidebar/thread-button";

export default function ThreadButtons() {
  const [threads, setThreads] = useState(Array(6).fill({name: 'Some Thread'}));

  return (
    <div className="h-full overflow-y-auto">
      {
        threads.map((thread, i) => {
          return (
            <div className="pl-2 pr-2">
              <ThreadButton thread={thread} />
            </div>
          )
        })
      }
    </div>
  )

}
