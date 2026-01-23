package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"baycommit-go/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client     *ethclient.Client //RPC 연결 객체
	chainID    *big.Int
	privateKey *ecdsa.PrivateKey
)

// 블록체인 클라이언트 초기화 (main에서 호출)
func ConnectBlockchain() error {
	var err error

	// RPC 연결
	rpcURL := os.Getenv("AVAX_RPC_HTTP")
	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("RPC 연결 실패: %v", err)
	}

	// Chain ID
	chainIDStr := os.Getenv("CHAIN_ID")
	chainID, _ = new(big.Int).SetString(chainIDStr, 10)

	// Admin Private Key (트랜잭션 서명용)
	pkHex := os.Getenv("ADMIN_PRIVATE_KEY")
	privateKey, err = crypto.HexToECDSA(pkHex)
	if err != nil {
		return fmt.Errorf("프라이빗 키 파싱 실패: %v", err)
	}

	return nil
}

// 컨트랙트 인스턴스 생성 (내부용)
func getContractInstance(proxyAddr string) (*contract.StudyGroup, error) {
	addr := common.HexToAddress(proxyAddr)
	return contract.NewStudyGroup(addr, client)
}

// 트랜잭션 옵션 생성 (쓰기용)
func getTxOpts() (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// ============ 읽기 함수 ============

// 컨트랙트 조회: 오늘 스터디 정보
func GetStudyDayInfo(proxyAddr string, timestamp uint64) (participantCount uint64, isClosed bool, err error) {
	instance, err := getContractInstance(proxyAddr)
	if err != nil {
		return 0, false, err
	}

	result, err := instance.GetStudyDayInfo(nil, new(big.Int).SetUint64(timestamp))
	if err != nil {
		return 0, false, err
	}

	return result.ParticipantCount.Uint64(), result.IsClosed, nil
}

// 컨트랙트 조회: 유저 커밋 시간
func GetCommitTime(proxyAddr string, timestamp uint64, userAddr string) (uint64, error) {
	instance, err := getContractInstance(proxyAddr)
	if err != nil {
		return 0, err
	}

	result, err := instance.GetCommitTime(nil, new(big.Int).SetUint64(timestamp), common.HexToAddress(userAddr))
	if err != nil {
		return 0, err
	}

	return result.Uint64(), nil
}

// ============ 쓰기 함수 ============

// 컨트랙트 호출: 오늘 스터디 시작
func StartTodayStudy(proxyAddr string, timestamp uint64) error {
	instance, err := getContractInstance(proxyAddr)
	if err != nil {
		return err
	}

	auth, err := getTxOpts()
	if err != nil {
		return err
	}

	tx, err := instance.StartTodayStudy(auth, new(big.Int).SetUint64(timestamp))
	if err != nil {
		return fmt.Errorf("StartTodayStudy 실패: %v", err)
	}

	fmt.Printf("StartTodayStudy tx 전송됨: %s\n", tx.Hash().Hex())

	// 트랜잭션 확인 대기
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("트랜잭션 마이닝 대기 실패: %v", err)
	}

	return nil
}

// 컨트랙트 호출: 커밋 기록
func TrackCommit(proxyAddr string, timestamp uint64, userAddr string, commitTime uint64) error {
	instance, err := getContractInstance(proxyAddr)
	if err != nil {
		return err
	}

	auth, err := getTxOpts()
	if err != nil {
		return err
	}

	tx, err := instance.TrackCommit(
		auth,
		new(big.Int).SetUint64(timestamp),
		common.HexToAddress(userAddr),
		new(big.Int).SetUint64(commitTime),
	)
	if err != nil {
		return fmt.Errorf("TrackCommit 실패: %v", err)
	}

	fmt.Printf("TrackCommit tx 전송됨: %s\n", tx.Hash().Hex())

	// 트랜잭션 확인 대기
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("트랜잭션 마이닝 대기 실패: %v", err)
	}

	return nil
}
