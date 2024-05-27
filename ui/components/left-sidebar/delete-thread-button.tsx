"use client"

import {toast} from 'sonner';
import {graphql, useMutation} from "react-relay"
import {ThreadButtonsQuery} from "@/components/left-sidebar/ThreadButtons";

interface props {
    id: string
}

const DeleteThreadButtonMutation = graphql`
mutation deleteThreadButtonMutation($id:ID!){
  deleteThread(id:$id)
}
`

export default function DeleteThreadButton({id}: props) {
    const [commitMutation, isMutationInFlight] = useMutation(DeleteThreadButtonMutation)
    return (
        <div
            className="flex flex-row items-center hover:bg-default-200 cursor-pointer pl-4 pr-4 pt-2 pb-2"
            onClick={() => {
                commitMutation({
                    variables: {
                        id: id,
                    },
                    onError: (e => {
                        toast.error('Error deleting thread')
                    }),
                    onCompleted: ((response) => {
                        toast.success('Deleted thread')
                    }),
                    optimisticUpdater: store => {
                        const {updatableData} = store.readUpdatableQuery(ThreadButtonsQuery, {})
                        // @ts-ignore
                        const viewer = updatableData.viewer
                        if (viewer != null) {
                            console.log('weh')
                        }
                    }
                })
            }}>
            <div className="mr-4">
                <img src="/trash.svg" alt="delete thread"/>
            </div>
            Delete
        </div>
    )
}