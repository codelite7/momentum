"use client"

import {
    Card,
    CardBody,
    Input,
    Modal,
    ModalBody,
    ModalContent,
    ModalHeader,
    Spinner
} from "@nextui-org/react";
import useSearchModalStore from "@/stores/search-modal-store";
import {useShallow} from "zustand/react/shallow";
import {useDebouncedCallback} from 'use-debounce';
import {Suspense, useState} from "react";
import SearchResults from "@/components/search/search-results";
import dynamic from "next/dynamic";
import MessageResult from "@/components/search/message-result";

const DynamicSearchResults = dynamic(() => import('@/components/search/search-results'), {
    loading: () => {
        return (
            <div className="flex place-content-center">
                <Spinner size="lg" color="primary" />
            </div>
        )
    },
})

export default function SearchModal() {
    const [isOpen, toggle] = useSearchModalStore(useShallow((state: any) => [state.isOpen, state.toggle]));
    const [matchedMessages, setMatchedMessages] = useState([])
    const [query, setQuery] = useState("");
    const debounced = useDebouncedCallback((query) => {
        setQuery(query);
    }, 1000);
    return (
        <Modal isOpen={isOpen} onClose={() => {
            debounced.cancel()
            setQuery('')
            toggle()
        }} size="4xl" classNames={{
            body: "px-0 mt-4",
            header: "px-0 mt-4 border-b-1",
        }}>
            <ModalContent>
                <ModalHeader>
                    <div className="flex w-full items-center p-4">
                        <Input
                            autoFocus
                            onValueChange={debounced}
                            size="sm"
                            color="primary"
                            variant="bordered"
                        />
                    </div>
                </ModalHeader>
                <ModalBody>
                    <Card>
                        <CardBody>
                            {
                                query ? (
                                    <DynamicSearchResults query={query}/>
                                ) : (
                                    <div className="flex w-full place-content-center">
                                        <span>Type to search <span className="tracking-widest">STRATIFI</span></span>
                                    </div>
                                )
                            }
                        </CardBody>
                    </Card>

                </ModalBody>
            </ModalContent>
        </Modal>
    )
}