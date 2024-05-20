import { EventEmitter, inject, Injectable } from '@angular/core'
import { GraphqlService, OrderBy } from './graphql.service'
import { AuthService } from '../auth.service'
import { gql } from 'graphql-request'
import {
  BookmarkFragment,
  BookmarkQuery,
  BookmarksQuery,
  BookmarksQueryVariables, CreateBookmarkMutation, CreateBookmarkMutationVariables,
  CreateThreadMutation, DeleteBookmarkMutation, MessageFragment,
  ThreadQuery,
  ThreadsQuery,
  ThreadsQueryVariables
} from '../../../graphql/generated'

@Injectable({
  providedIn: 'root'
})
export class BookmarkService {
  graphqlService: GraphqlService = inject(GraphqlService)
  authService: AuthService = inject(AuthService)
  bookmarkCreatedEmitter: EventEmitter<string> = new EventEmitter()
  bookmarkUpdatedEmitter: EventEmitter<string> = new EventEmitter()
  bookmarkDeletedEmitter: EventEmitter<string> = new EventEmitter()

  constructor() { }
  async bookmarks(variables?: BookmarksQueryVariables): Promise<BookmarkFragment[]> {
    let response = await this.graphqlService.sdk.bookmarks(variables)
    if (response.bookmarks.edges) {
      let bookmarks: BookmarkFragment[] = []
      response.bookmarks.edges.forEach(edge => bookmarks.push(edge?.node as BookmarkFragment))
      return bookmarks
    }
    return []
  }

  async bookmark(id:string): Promise<BookmarkFragment | undefined> {
    let response = await this.graphqlService.sdk.bookmark({id:id})
    if (response.bookmarks.edges) {
      return response.bookmarks.edges[0]?.node as BookmarkFragment
    }
    return undefined
  }

  async bookmarkMessage(id: string): Promise<BookmarkFragment> {
    let userId = await this.authService.userId()
    let vars: CreateBookmarkMutationVariables = {input: {messageID: id, userID: userId}}
    let response = await this.graphqlService.sdk.createBookmark(vars)
    this.bookmarkCreatedEmitter.emit(response.createBookmark?.id)
    return response.createBookmark as BookmarkFragment
  }

  async bookmarkResponse(id: string): Promise<BookmarkFragment> {
    let userId = await this.authService.userId()
    let vars: CreateBookmarkMutationVariables = {input: {responseID: id, userID: userId}}
    let response = await this.graphqlService.sdk.createBookmark(vars)
    this.bookmarkCreatedEmitter.emit(response.createBookmark?.id)
    return response.createBookmark as BookmarkFragment
  }

  async bookmarkThread(id: string): Promise<BookmarkFragment> {
    let userId = await this.authService.userId()
    let vars: CreateBookmarkMutationVariables = {input: {threadID: id, userID: userId}}
    let response = await this.graphqlService.sdk.createBookmark(vars)
    this.bookmarkCreatedEmitter.emit(response.createBookmark?.id)
    return response.createBookmark as BookmarkFragment
  }

  async deleteBookmark(id: string): Promise<void> {
    await this.graphqlService.sdk.deleteBookmark({id:id})
    this.bookmarkDeletedEmitter.emit(id)
  }
}
