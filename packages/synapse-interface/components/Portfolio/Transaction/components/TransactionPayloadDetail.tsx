import { useCallback } from 'react'
import Image from 'next/image'
import { useAppDispatch } from '@/store/hooks'
import {
  setFromChainId,
  setToChainId,
  setFromToken,
  setToToken,
} from '@/slices/bridge/reducer'
import { Chain, Token } from '@/utils/types'

export const TransactionPayloadDetail = ({
  chain,
  token,
  tokenAmount,
  isOrigin,
}: {
  chain?: Chain
  token?: Token
  tokenAmount?: number
  isOrigin: boolean
}) => {
  const dispatch = useAppDispatch()

  const handleSelectChainCallback = useCallback(() => {
    if (isOrigin) {
      dispatch(setFromChainId(chain?.id as number))
    } else {
      dispatch(setToChainId(chain?.id as number))
    }
  }, [isOrigin, chain])

  const handleSelectTokenCallback = useCallback(() => {
    if (isOrigin && chain && token) {
      dispatch(setFromChainId(chain?.id as number))
      dispatch(setFromToken(token as Token))
    } else {
      dispatch(setToChainId(chain?.id as number))
      dispatch(setToToken(token as Token))
    }
  }, [isOrigin, token, chain])

  return (
    <div
      data-test-id="transaction-payload-detail"
      className="flex flex-col p-1 space-y-1"
    >
      {chain && (
        <div
          data-test-id="transaction-payload-network"
          onClick={handleSelectChainCallback}
          className={`
            flex flex-row px-1 items-center cursor-pointer rounded-sm w-fit
            hover:bg-tint active:opacity-[67%]
          `}
        >
          <Image
            src={chain.chainImg}
            className="w-4 h-4 mr-1.5 rounded-full"
            alt={`${chain.name} icon`}
          />
          <div>{chain.name}</div>
        </div>
      )}

      {token && tokenAmount && (
        <div
          data-test-id="transaction-payload-token"
          onClick={handleSelectTokenCallback}
          className={`
            flex flex-row px-1 items-center cursor-pointer rounded-sm w-fit
            hover:bg-tint active:opacity-[67%]
          `}
        >
          <Image
            src={token?.icon}
            className="items-center w-4 h-4 mr-1.5 rounded-full"
            alt={`${token?.name} icon`}
          />
          {typeof tokenAmount === 'number' ? (
            <div className="mr-1">{tokenAmount}</div>
          ) : (
            <div className="mr-1">...</div>
          )}
          <div className="mt-0.5 text-sm">{token?.symbol}</div>
        </div>
      )}
    </div>
  )
}
