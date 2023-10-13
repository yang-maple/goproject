package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type secret struct{}
type Secretsresp struct {
	Total int             `json:"total"`
	Item  []corev1.Secret `json:"item"`
}

type Createsecret struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
	Immutable bool              `json:"immutable"`
	Type      string            `json:"type"`
	Data      map[string]string `json:"data"`
}

var Secrets secret

func (s secret) tocells(secrets []corev1.Secret) []DataCell {
	cells := make([]DataCell, len(secrets))
	for i := range secrets {
		cells[i] = secretCell(secrets[i])
	}
	return cells
}

func (s secret) fromcells(cells []DataCell) []corev1.Secret {
	secrets := make([]corev1.Secret, len(cells))
	for i := range cells {
		secrets[i] = corev1.Secret(cells[i].(secretCell))
	}
	return secrets
}

// 列表
func (s *secret) GetSecretList(Secretname, Namespace string, Limit, Page int) (DP *Secretsresp, err error) {
	//获取deployment 的所有清单列表
	secretlist, err := K8s.Clientset.CoreV1().Secrets(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取 secret 失败" + err.Error())
	}

	//组装数据
	selectdata := &dataselector{
		GenericDataList: s.tocells(secretlist.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Secretname},
			Paginate: &PaginateQuery{
				limit: Limit,
				page:  Page,
			},
		},
	}
	//先过滤 后排序
	filtered := selectdata.Filter()
	total := len(filtered.GenericDataList)
	//分页
	datapage := filtered.Sort().Pagination()
	secrets := s.fromcells(datapage.GenericDataList)
	for _, v := range secrets {
		fmt.Println(v.Name, v.CreationTimestamp.Time)
	}
	return &Secretsresp{
		Total: total,
		Item:  secrets,
	}, err
}

// 详情
func (s *secret) GetSecretDetal(Namespace, Secretname string) (detail *corev1.Secret, err error) {
	//获取deploy
	detail, err = K8s.Clientset.CoreV1().Secrets(Namespace).Get(context.TODO(), Secretname, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取Secret 详情失败" + err.Error())
		return nil, errors.New("获取Secret 详情失败" + err.Error())
	}
	return detail, nil
}

// 创建
func (s *secret) CreateSecret(data *Createsecret) (err error) {

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:   data.Name,
			Labels: data.Labels,
		},
		Immutable:  &data.Immutable,
		Data:       nil,
		StringData: data.Data,
		Type:       corev1.SecretType(data.Type),
	}

	_, err = K8s.Clientset.CoreV1().Secrets(data.Namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		logger.Info("创建 secret 失败" + err.Error())
		return errors.New("创建 secret 失败" + err.Error())
	}
	return nil

}

// 删除
func (s *secret) DelSecret(Namespace, Secretname string) (err error) {
	err = K8s.Clientset.CoreV1().Secrets(Namespace).Delete(context.TODO(), Secretname, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除Secret失败" + err.Error())
		return errors.New("删除Secret失败" + err.Error())
	}
	return nil
}
