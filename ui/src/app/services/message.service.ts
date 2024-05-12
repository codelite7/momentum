import { EventEmitter, inject, Injectable } from '@angular/core'
import { gql, GraphQLClient } from 'graphql-request'
import { GraphqlService, ListArgs, OrderBy } from './graphql.service'
import { AuthService } from '../auth.service'
import {
  CreateMessageMutation,
  CreateThreadMutation,
  MessageFragment,
  MessageQuery,
  MessagesQuery,
  MessagesQueryVariables,
  ThreadFragment,
  ThreadMessageFragment,
  ThreadMessageQueryVariables,
  ThreadMessagesQuery,
  ThreadMessagesQueryVariables,
  ThreadQuery,
  ThreadsQuery,
  ThreadsQueryVariables
} from '../../../graphql/generated'

@Injectable({
  providedIn: 'root'
})
export class MessageService {
  graphqlService: GraphqlService = inject(GraphqlService)
  authService: AuthService = inject(AuthService)
  messageCreatedEmitter: EventEmitter<string> = new EventEmitter()
  messageUpdatedEmitter: EventEmitter<string> = new EventEmitter()
  messageDeletedEmitter: EventEmitter<string> = new EventEmitter()
  constructor() { }

  async messages(variables?: MessagesQueryVariables): Promise<MessageFragment[]> {
    let response = await this.graphqlService.sdk.messages(variables)
    if (response.messages.edges) {
      let messages: MessageFragment[] = []
      response.messages.edges.forEach(edge => messages.push(edge?.node as MessageFragment))
      return messages
    }
    return []
  }

  async message(id:string): Promise<MessageFragment | undefined> {
    let response = await this.graphqlService.sdk.message({id:id})
    if (response.messages.edges) {
      return response.messages.edges[0] as MessageFragment
    }
    return undefined
  }

  async createMessage(content: string, threadID: string): Promise<MessageFragment> {
    let userId = await this.authService.userId()
    let response = await this.graphqlService.sdk.createMessage({input: {content: content, sentByUserID: userId, threadID: threadID}})
    this.messageCreatedEmitter.emit(response.createMessage?.id)
    return response.createMessage as MessageFragment
  }

  async threadMessages(vars: ThreadMessagesQueryVariables): Promise<ThreadMessageFragment[]> {
    vars.userId = await this.authService.userId()
    let response =  await this.graphqlService.sdk.threadMessages(vars)
    if (response.messages.edges) {
      let messages: ThreadMessageFragment[] = []
      response.messages.edges.forEach(edge => messages.push(edge?.node as ThreadMessageFragment))
      return messages
    }
    return []
  }

  async threadMessage(vars: ThreadMessageQueryVariables): Promise<ThreadMessageFragment | undefined> {
    vars.userId = await this.authService.userId()
    let response =  await this.graphqlService.sdk.threadMessage(vars)
    if (response.messages.edges && response.messages.edges.length == 1) {
      let message = response.messages.edges.pop()
      if (message && message.node) {
        return message.node as ThreadMessageFragment
      }
    }
    return undefined
  }
}
