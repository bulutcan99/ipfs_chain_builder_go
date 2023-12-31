package main

import (
	"encoding/json"
	"fmt"
	"github.com/bulutcan99/go_ipfs_chain_builder/internal/repository"
	"github.com/bulutcan99/go_ipfs_chain_builder/internal/service"
	"github.com/bulutcan99/go_ipfs_chain_builder/ipfs"
	"github.com/bulutcan99/go_ipfs_chain_builder/model"
	config_mysql "github.com/bulutcan99/go_ipfs_chain_builder/pkg/config/mysql"
	"github.com/bulutcan99/go_ipfs_chain_builder/pkg/env"
	"github.com/bulutcan99/go_ipfs_chain_builder/pkg/logger"
	shell "github.com/ipfs/go-ipfs-api"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

var (
	MYSQL  *config_mysql.MYSQL
	Env    *env.ENV
	Logger *zap.Logger
)

func Init() {
	Env = env.ParseEnv()
	MYSQL = config_mysql.NewMYSQLConnection()
	Logger = logger.InitLogger(Env.LogLevel)
}

func main() {
	Init()
	ipfsCon := fmt.Sprintf("localhost:%d", Env.IpfsPort)
	sh := shell.NewShell(ipfsCon)
	var firstNode, prevNode *ipfs.Node
	defer MYSQL.Close()
	defer Logger.Sync()
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
		PRIMARY KEY (id)
		);
	`

	_, err := MYSQL.Client.Exec(createTableQuery)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("users table created successfully")
	userRepo := repository.NewUserRepo(MYSQL)
	userService := service.NewUserService(userRepo)

	users := []model.User{
		{
			Name: "Bulut",
		},
		{
			Name: "Can",
		},
		{
			Name: "Gocer",
		},
	}

	for _, user := range users {
		_, err := userService.AddUser(user)
		if err != nil {
			zap.S().Error("ERROR:", err)
		}

	}
	aggregateRepo := repository.NewAggregateRepo(MYSQL)
	aggregateService := service.NewAggregateService(aggregateRepo)
	aggregatedData, err := aggregateService.GetUsersWithColumnTypes()
	if err != nil {
		zap.S().Error("ERROR:", err)
	}

	for _, d := range aggregatedData {
		var prevHash string
		if prevNode != nil {
			prevHash = prevNode.Hash
		}

		currentNode := ipfs.NewNode(d, prevHash)

		if prevNode == nil {
			firstNode = currentNode
		} else {
			prevNode.Next = currentNode
		}

		prevNode = currentNode
	}

	chainBytes, err := json.Marshal(firstNode)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error: %s", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	cid, err := sh.Add(strings.NewReader(string(chainBytes)))
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error: %s", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	fmt.Printf("Added hash chain to IPFS with CID: %s\n", cid)

	response, err := sh.PublishWithDetails(cid, "", time.Hour, time.Hour, false)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error: %s", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	fmt.Printf("Published to IPNS with key:", response)
}
