import { createAction } from '@reduxjs/toolkit'

import { Chain, Token } from '@/utils/types'

export interface PendingBridgeTransaction {
  id: number
  originChain: Chain
  originToken: Token
  originValue: string
  destinationChain: Chain
  destinationToken: Token
  transactionHash?: string
  timestamp: number
  isSubmitted: boolean
}

export const addPendingBridgeTransaction =
  createAction<PendingBridgeTransaction>('bridge/addPendingBridgeTransaction')
export const updatePendingBridgeTransaction = createAction<{
  id: number
  timestamp: number
  transactionHash: string
  isSubmitted: boolean
}>('bridge/updatePendingBridgeTransaction')
export const removePendingBridgeTransaction = createAction<number>(
  'bridge/removePendingBridgeTransaction'
)
export const updatePendingBridgeTransactions = createAction<
  PendingBridgeTransaction[]
>('bridge/updatePendingBridgeTransactions')
