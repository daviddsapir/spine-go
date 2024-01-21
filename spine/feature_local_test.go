package spine

import (
	"testing"

	"github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/mocks"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestDeviceClassificationSuite(t *testing.T) {
	suite.Run(t, new(DeviceClassificationTestSuite))
}

type DeviceClassificationTestSuite struct {
	suite.Suite
	senderMock                                                *mocks.Sender
	localDevice                                               *DeviceLocalImpl
	localEntity                                               *EntityLocalImpl
	function, serverWriteFunction                             model.FunctionType
	featureType, subFeatureType                               model.FeatureTypeType
	msgCounter                                                model.MsgCounterType
	remoteFeature, remoteServerFeature, remoteSubFeature      api.FeatureRemote
	localFeature, localServerFeature, localServerFeatureWrite api.FeatureLocal
}

func (suite *DeviceClassificationTestSuite) BeforeTest(suiteName, testName string) {
	suite.senderMock = mocks.NewSender(suite.T())
	suite.function = model.FunctionTypeDeviceClassificationManufacturerData
	suite.featureType = model.FeatureTypeTypeDeviceClassification
	suite.subFeatureType = model.FeatureTypeTypeLoadControl
	suite.serverWriteFunction = model.FunctionTypeLoadControlLimitListData
	suite.msgCounter = model.MsgCounterType(1)

	suite.localDevice, suite.localEntity = createLocalDeviceAndEntity(1)
	suite.localFeature, suite.localServerFeature = createLocalFeatures(suite.localDevice, suite.localEntity, suite.featureType, "")
	_, suite.localServerFeatureWrite = createLocalFeatures(suite.localDevice, suite.localEntity, suite.subFeatureType, suite.serverWriteFunction)

	remoteDevice := createRemoteDevice(suite.localDevice, suite.senderMock)
	suite.remoteFeature, suite.remoteServerFeature = createRemoteEntityAndFeature(suite.localDevice, remoteDevice, 1, suite.featureType)
	suite.remoteSubFeature, _ = createRemoteEntityAndFeature(suite.localDevice, remoteDevice, 2, suite.subFeatureType)
}

func (suite *DeviceClassificationTestSuite) TestDeviceClassification_Request_Reply() {
	suite.senderMock.On("Request", model.CmdClassifierTypeRead, suite.localFeature.Address(), suite.remoteFeature.Address(), false, mock.AnythingOfType("[]model.CmdType")).Return(&suite.msgCounter, nil)

	// send data request
	msgCounter, err := suite.localFeature.RequestData(suite.function, nil, nil, suite.remoteFeature)
	assert.Nil(suite.T(), err)

	manufacturerData := &model.DeviceClassificationManufacturerDataType{
		BrandName:    util.Ptr(model.DeviceClassificationStringType("brand name")),
		VendorName:   util.Ptr(model.DeviceClassificationStringType("vendor name")),
		DeviceName:   util.Ptr(model.DeviceClassificationStringType("device name")),
		DeviceCode:   util.Ptr(model.DeviceClassificationStringType("device code")),
		SerialNumber: util.Ptr(model.DeviceClassificationStringType("serial number")),
	}

	replyMsg := api.Message{
		Cmd: model.CmdType{
			DeviceClassificationManufacturerData: manufacturerData,
		},
		CmdClassifier: model.CmdClassifierTypeReply,
		RequestHeader: &model.HeaderType{
			MsgCounter:          util.Ptr(model.MsgCounterType(1)),
			MsgCounterReference: &suite.msgCounter,
		},
		FeatureRemote: suite.remoteFeature,
	}
	// set response
	msgErr := suite.localFeature.HandleMessage(&replyMsg)
	if assert.Nil(suite.T(), msgErr) {
		remoteData := suite.remoteFeature.DataCopy(suite.function)
		assert.IsType(suite.T(), &model.DeviceClassificationManufacturerDataType{}, remoteData, "Data has wrong type")
	}

	// Act
	result, err := suite.localFeature.FetchRequestData(*msgCounter, suite.remoteFeature)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.IsType(suite.T(), &model.DeviceClassificationManufacturerDataType{}, result, "Data has wrong type")
	receivedData := result.(*model.DeviceClassificationManufacturerDataType)

	assert.Equal(suite.T(), manufacturerData.BrandName, receivedData.BrandName)
	assert.Equal(suite.T(), manufacturerData.VendorName, receivedData.VendorName)
	assert.Equal(suite.T(), manufacturerData.DeviceName, receivedData.DeviceName)
	assert.Equal(suite.T(), manufacturerData.DeviceCode, receivedData.DeviceCode)
	assert.Equal(suite.T(), manufacturerData.SerialNumber, receivedData.SerialNumber)
}

func (suite *DeviceClassificationTestSuite) TestDeviceClassification_Request_Error() {
	suite.senderMock.On("Request", model.CmdClassifierTypeRead, suite.localFeature.Address(), suite.remoteFeature.Address(), false, mock.AnythingOfType("[]model.CmdType")).Return(&suite.msgCounter, nil)

	const errorNumber = model.ErrorNumberTypeGeneralError
	const errorDescription = "error occurred"

	// send data request
	msgCounter, err := suite.localFeature.RequestData(suite.function, nil, nil, suite.remoteFeature)
	assert.Nil(suite.T(), err)

	replyMsg := api.Message{
		Cmd: model.CmdType{
			ResultData: &model.ResultDataType{
				ErrorNumber: util.Ptr(model.ErrorNumberType(errorNumber)),
				Description: util.Ptr(model.DescriptionType(errorDescription)),
			},
		},
		CmdClassifier: model.CmdClassifierTypeResult,
		RequestHeader: &model.HeaderType{
			MsgCounter:          util.Ptr(model.MsgCounterType(1)),
			MsgCounterReference: &suite.msgCounter,
		},
		FeatureRemote: suite.remoteFeature,
		EntityRemote:  suite.remoteFeature.Entity(),
		DeviceRemote:  suite.remoteFeature.Device(),
	}

	// set response
	msgErr := suite.localFeature.HandleMessage(&replyMsg)
	if assert.Nil(suite.T(), msgErr) {
		remoteData := suite.remoteFeature.DataCopy(suite.function)
		assert.Nil(suite.T(), remoteData)
	}

	// Act
	result, err := suite.localFeature.FetchRequestData(*msgCounter, suite.remoteFeature)
	assert.Nil(suite.T(), result)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errorNumber, err.ErrorNumber)
	assert.Equal(suite.T(), errorDescription, string(*err.Description))
}

func (suite *DeviceClassificationTestSuite) TestDeviceClassification_Subscribiptions() {
	suite.senderMock.On("Subscribe", mock.Anything, mock.Anything, mock.Anything).Return(&suite.msgCounter, nil)
	suite.senderMock.On("Unsubscribe", mock.Anything, mock.Anything, mock.Anything).Return(&suite.msgCounter, nil)

	msgCounter, err := suite.localFeature.Subscribe(suite.remoteFeature.Address())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), msgCounter)

	suite.localFeature.RemoveSubscription(suite.remoteFeature.Address())

	suite.localFeature.Device().AddRemoteDeviceForSki(suite.remoteFeature.Device().Ski(), suite.remoteFeature.Device())

	msgCounter, err = suite.localServerFeature.Subscribe(suite.remoteFeature.Address())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), msgCounter)

	suite.localFeature.RemoveSubscription(suite.remoteFeature.Address())

	msgCounter, err = suite.localFeature.Subscribe(suite.remoteFeature.Address())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), msgCounter)

	msgCounter, err = suite.localFeature.Subscribe(suite.remoteSubFeature.Address())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), msgCounter)

	suite.localFeature.RemoveSubscription(suite.remoteFeature.Address())

	suite.localFeature.RemoveAllSubscriptions()
}

func (suite *DeviceClassificationTestSuite) TestDeviceClassification_Bindings() {
	suite.senderMock.On("Bind", mock.Anything, mock.Anything, mock.Anything).Return(&suite.msgCounter, nil)
	suite.senderMock.On("Unbind", mock.Anything, mock.Anything, mock.Anything).Return(&suite.msgCounter, nil)

	msgCounter, err := suite.localFeature.Bind(suite.remoteFeature.Address())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), msgCounter)

	suite.localFeature.RemoveBinding(suite.remoteFeature.Address())

	suite.localFeature.Device().AddRemoteDeviceForSki(suite.remoteFeature.Device().Ski(), suite.remoteFeature.Device())

	msgCounter, err = suite.localServerFeature.Bind(suite.remoteFeature.Address())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), msgCounter)

	suite.localFeature.RemoveBinding(suite.remoteFeature.Address())

	msgCounter, err = suite.localFeature.Bind(suite.remoteFeature.Address())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), msgCounter)

	msgCounter, err = suite.localFeature.Bind(suite.remoteSubFeature.Address())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), msgCounter)

	suite.localFeature.RemoveBinding(suite.remoteFeature.Address())

	suite.localFeature.RemoveAllBindings()
}

func (suite *DeviceClassificationTestSuite) Test_Reply() {
	msg := &api.Message{
		FeatureRemote: suite.remoteServerFeature,
		CmdClassifier: model.CmdClassifierTypeReply,
		Cmd: model.CmdType{
			DeviceClassificationManufacturerData: &model.DeviceClassificationManufacturerDataType{},
		},
	}

	err := suite.localFeature.HandleMessage(msg)
	assert.Nil(suite.T(), err)

	msg.RequestHeader = &model.HeaderType{
		MsgCounterReference: util.Ptr(model.MsgCounterType(100)),
	}
	err = suite.localFeature.HandleMessage(msg)
	assert.Nil(suite.T(), err)
}

func (suite *DeviceClassificationTestSuite) Test_Notify() {
	msg := &api.Message{
		FeatureRemote: suite.remoteServerFeature,
		CmdClassifier: model.CmdClassifierTypeNotify,
		Cmd: model.CmdType{
			DeviceClassificationManufacturerData: &model.DeviceClassificationManufacturerDataType{},
		},
	}

	err := suite.localFeature.HandleMessage(msg)
	assert.Nil(suite.T(), err)
}

func (suite *DeviceClassificationTestSuite) Test_Write() {
	msg := &api.Message{
		CmdClassifier: model.CmdClassifierTypeWrite,
		FeatureRemote: suite.remoteSubFeature,
		Cmd: model.CmdType{
			LoadControlLimitListData: &model.LoadControlLimitListDataType{},
		},
	}

	err := suite.localServerFeatureWrite.HandleMessage(msg)
	assert.Nil(suite.T(), err)
}
