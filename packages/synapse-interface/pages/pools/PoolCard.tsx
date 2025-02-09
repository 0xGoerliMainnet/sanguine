import Link from 'next/link'
import { useRouter } from 'next/router'
import { getPoolUrl } from '@urls'
import { useEffect, useState } from 'react'
import { getSinglePoolData } from '@utils/actions/getPoolData'
import { getPoolApyData } from '@utils/actions/getPoolApyData'
import { Token } from '@types'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import { memo } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import LoadingSpinner from '@tw/LoadingSpinner'
import { useAccount } from 'wagmi'
import { toast } from 'react-hot-toast'
import { commify, formatBigIntToString } from '@/utils/bigint/format'
import { stringToBigInt } from '@/utils/bigint/format'

const PoolCard = memo(
  ({
    pool,
    chainId,
    address,
    connectedChainId,
    prices,
  }: {
    pool: Token
    chainId: number
    address: string
    connectedChainId: number
    prices: any
  }) => {
    const [poolData, setPoolData] = useState(undefined)
    const [poolApyData, setPoolApyData] = useState(undefined)
    const { isDisconnected } = useAccount()
    let popup: string

    useEffect(() => {
      if (connectedChainId && chainId && pool) {
        // TODO - separate the apy and tvl so they load async.
        getSinglePoolData(chainId, pool, prices)
          .then((res) => {
            setPoolData(res)
          })
          .catch((err) => {
            console.log('Could not get Pool Data: ', err)
          })
        getPoolApyData(chainId, pool, prices)
          .then((res) => {
            setPoolApyData(res)
          })
          .catch((err) => {
            console.log('Could not get Pool APY Data: ', err)
          })
      }
    }, [])
    const chain = CHAINS_BY_ID[chainId]

    /*
  useEffect triggers: address, isDisconnected, popup
  - will dismiss toast asking user to connect wallet once wallet has been connected
  */
    useEffect(() => {
      if (address && !isDisconnected && popup) {
        toast.dismiss(popup)
      }
    }, [address, isDisconnected, popup])

    return (
      <div>
        <Link
          onClick={() => {
            if (address === undefined || isDisconnected) {
              popup = toast.error('Please connect your wallet', {
                id: 'pools-connect-wallet',
                duration: 20000,
              })
              return popup
            }
          }}
          href={getPoolUrl(pool)}
        >
          <Card
            title={
              <PoolsCardTitle
                chainImg={chain?.chainImg?.src}
                poolName={pool?.poolName}
                chainName={chain?.name}
              />
            }
            titleClassName="text-white font-light text-xl"
            className={`
            bg-bgBase transition-all rounded-md items-center
            hover:bg-bgLight
              py-6 mt-4
              border border-transparent
              whitespace-wrap
            `}
            divider={false}
          >
            <Grid gap={3} cols={{ xs: 3 }} className="pt-8">
              <div>
                <h3 className="text-sm text-opacity-50 text-secondaryTextColor">
                  Assets
                </h3>
                {pool?.poolTokens && (
                  <CoinLabels coins={pool.poolTokens} /> // change coin to token
                )}
              </div>
              <div>
                <h3 className="text-sm text-opacity-50 text-secondaryTextColor">
                  TVL
                </h3>
                <div className={'mt-2 text-white '}>
                  {poolData?.totalLockedUSD ? (
                    `$${commify(
                      formatBigIntToString(
                        stringToBigInt(
                          `${poolData.totalLockedUSD}`,
                          pool.decimals[chainId]
                        ),
                        18,
                        -1
                      )
                    )}`
                  ) : (
                    <LoadingSpinner shift={true} />
                  )}
                </div>
              </div>
              <div>
                <h3 className="text-sm text-opacity-50 text-secondaryTextColor">
                  APY{' '}
                </h3>
                <div className="mt-2 text-green-400">
                  <>
                    {poolApyData?.fullCompoundedAPYStr
                      ? String(poolApyData.fullCompoundedAPYStr)
                      : '-'}
                    %
                  </>
                </div>
              </div>
            </Grid>
          </Card>
        </Link>
      </div>
    )
  }
)

const PoolsCardTitle = ({
  chainName,
  poolName,
  chainImg,
}: {
  chainName: string
  poolName: string
  chainImg: string
}) => {
  let displayPoolName = poolName?.replace(chainName, `${chainName}`)

  return (
    <div className="flex items-center">
      <img src={chainImg} className="w-6 h-6 mr-2 rounded-full" />
      <div className="font-semibold">{displayPoolName}</div>
    </div>
  )
}

const CoinLabels = ({ coins }) => {
  return (
    <div className="flex mt-3">
      {coins
        ? coins?.map((coin, i) => (
            <img
              alt=""
              className="w-5 mr-1 rounded-full"
              src={coin.icon.src}
              key={i}
            />
          ))
        : null}
    </div>
  )
}
export default PoolCard
