import { EventEmitter, inject, Injectable } from '@angular/core'
import { GraphqlService, ListArgs, OrderBy } from './graphql.service'
import { gql } from 'graphql-request'
import { AuthService } from '../auth.service'
import {
  CreateThreadMutation,
  getSdk, Thread, ThreadFragment,
  ThreadQuery,
  ThreadsQuery,
  ThreadsQueryVariables
} from '../../../graphql/generated'

@Injectable({
  providedIn: 'root'
})
export class ThreadService {
  graphqlService: GraphqlService = inject(GraphqlService)
  authService: AuthService = inject(AuthService)
  threadCreatedEmitter: EventEmitter<string> = new EventEmitter()
  threadUpdatedEmitter: EventEmitter<string> = new EventEmitter()
  threadDeletedEmitter: EventEmitter<string> = new EventEmitter()
  constructor() {
  }

  async threads(variables?: ThreadsQueryVariables): Promise<ThreadFragment[]> {
    let response = await this.graphqlService.sdk.threads(variables)
    if (response.threads.edges) {
      let threads: ThreadFragment[] = []
      response.threads.edges.forEach(edge => threads.push(edge?.node as ThreadFragment))
      return threads
    }
    return []
  }

  async thread(id:string): Promise<ThreadFragment | undefined> {
    let response = await this.graphqlService.sdk.thread({id:id})
    if (response.threads.edges) {
      return response.threads.edges[0]?.node as ThreadFragment
    }
    return undefined
  }

  async createThread(name:string, parentId?: string): Promise<ThreadFragment> {
    let userId = await this.authService.userId()
    let response = await this.graphqlService.sdk.createThread({input: {name:name, createdByID: userId, parentID: parentId}})
    this.threadCreatedEmitter.emit(response.createThread?.id)
    return response.createThread as ThreadFragment
  }
}
