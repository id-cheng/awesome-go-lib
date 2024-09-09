package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	// Load in a dataset, with headers. Header attributes will be stored.
	// Think of instances as a Data Frame structure in R or Pandas.
	// You can also create instances from scratch.
	// 加载数据集
	rawData, _ := base.ParseCSVToInstances("iris.csv", true)

	// Print a pleasant summary of your data.
	fmt.Println(rawData)

	// Initialises a new KNN classifier
	// 新的 KNN 分类器
	// "euclidean"：使用欧几里得距离作为距离度量。
	// "linear"：使用线性权重（即所有邻居对预测结果具有相同影响）。
	// 设置 K 值为 2，即考虑最近的两个邻居。
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	//Do a training-test split
	//将数据集按 50% 的比例分为训练集和测试集
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)

	//存储训练数据以供以后使用
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	//模型预测
	predictions, _ := cls.Predict(testData)

	// Prints precision/recall metrics
	confusionMat, _ := evaluation.GetConfusionMatrix(testData, predictions)
	fmt.Println(evaluation.GetSummary(confusionMat))
}
