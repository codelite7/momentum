import { Button } from "@nextui-org/button";
import {Card, CardHeader, CardBody, CardFooter} from "@nextui-org/card";
import {Popover, PopoverTrigger, PopoverContent} from "@nextui-org/popover";
interface ThreadButtonProps {
  thread: any
}
export default function ThreadButton({thread}: ThreadButtonProps) {
  return (
    <Button
      fullWidth
      size="sm"
      className="rounded pr-0 pl-2 mb-2 bg-forward-600"
    >
      <div className="w-full flex items-center justify-between">
        {thread.name}
        <Popover
          placement="bottom-start"
          classNames={{
            content: [
              "rounded border border-default-200 p-0 m-0"
            ]
          }}
        >
          <PopoverTrigger>
            <Card shadow="none" classNames={{base: "bg-forward-600"}}>
              <CardBody>
                <i className="pi pi-ellipsis-v"></i>
              </CardBody>
            </Card>
          </PopoverTrigger>
          <PopoverContent>
          <div className="flex flex-col">
              <div className="flex flex-row items-center hover:bg-default-200 cursor-pointer pl-4 pr-4 pt-2 pb-2">
                <div className="mr-4">
                  <img src="/edit.svg" alt="edit thread name"/>
                </div>
                Rename
              </div>
              <div className="flex flex-row items-center hover:bg-default-200 cursor-pointer pl-4 pr-4 pt-2 pb-2">
                <div className="mr-4">
                  <img src="/trash.svg" alt="delete thread"/>
                </div>
                Delete
              </div>
            </div>
          </PopoverContent>
        </Popover>
      </div>
    </Button>
  )
}
