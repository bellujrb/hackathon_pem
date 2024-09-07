import { Alert, ScrollView, StyleSheet, View } from "react-native";

import EditScreenInfo from "@/components/EditScreenInfo";
import { Text } from "@/components/Themed";
import { Button, ButtonText } from "@/components/ui/button";
import {
  Actionsheet,
  ActionsheetBackdrop,
  ActionsheetContent,
  ActionsheetDragIndicator,
  ActionsheetDragIndicatorWrapper,
  ActionsheetItem,
  ActionsheetItemText,
} from "@/components/ui/actionsheet";
import React from "react";
import { Divider } from "@/components/ui/divider";
import { Center } from "@/components/ui/center";
import { Heading } from "@/components/ui/heading";
import { Pressable } from "@/components/ui/pressable";

import { SafeAreaView } from "react-native-safe-area-context";
import { Linking } from 'react-native';
import { VStack } from "@/components/ui/vstack";


const CustomButton = ({
  iconName,
  label,
  color,
}: {
  iconName: string;
  label: string;
  color: string;
}) => {
  return (
    <Pressable
      onPress={() => alert(`Navegar para ${label}`)}
      style={{
        backgroundColor: "#FFFFFF",
        borderRadius: 10,
        padding: 16,
        justifyContent: "center",
        alignItems: "center",
        minWidth: 120,
        elevation: 3, // Leve sombra
        shadowColor: "#000",
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.3,
        shadowRadius: 2,
        marginTop: 10,
        marginBottom: 10,
      }}
    >
      <Text style={{ fontSize: 16, fontWeight: "bold", color }}>{label}</Text>
    </Pressable>
  );
};

export default function TabOneScreen() {
  const alertMessage = () => {
    Alert.alert(
      "Botão de Pânico",
      "Uma viatura sera encaminhada ao local, não se preocupe, temos a geolocalização do seu aparelho e saberemos onde você está. AGUARDE E SE MANTENHA EM SEGURANÇA!",
      [
        {
          text: "Cancelar",
          onPress: () => console.log("Cancel Pressed"),
          style: "cancel",
        },
        { text: "OK", onPress: () => console.log("OK Pressed") },
      ]
    );
  };

  const alertBoletim = () => {
    Alert.alert(
      "Boletim Online",
      "Você será redirecionado para o site da Policia Civil de Pernambuco. Deseja continuar?",
      [
        {
          text: "Cancelar",
          onPress: () => console.log("Cancel Pressed"),
          style: "cancel",
        },
        { text: "Sim", onPress: () => Linking.openURL('https://servicos.sds.pe.gov.br/delegacia/') },
      ]
    );
  }

  return (
    <SafeAreaView style={{ flex: 1, backgroundColor: "#f5f5f5" }}>
      <ScrollView>

        <View className="flex-1 justify-between px-4 h-full">
          <VStack space={6}>
            <View>
              <Button
                onPress={() => alertMessage()}
                style={{
                  backgroundColor: "#D32F2F",
                }}
              >
                <ButtonText>Botao de Panico</ButtonText>
              </Button>

              <Button
                onPress={() => alertBoletim()}
                style={{
                  backgroundColor: "#4CAF50",
                  marginTop: 10,
                }}
              >
                <ButtonText>Boletim Online</ButtonText>
              </Button>

            </View>
<Divider />
            <CustomButton
              iconName="gavel"
              label="Delegacias Especializadas no Atendimento à Mulher (DEAM)"
              color="#FF6B6B"
            />
            <CustomButton
              iconName="phone"
              label="Central de Atendimento à Mulher - Ligue 180"
              color="#4CAF50"
            />
            <CustomButton
              iconName="home"
              label="Casa Abrigo para Mulheres em Situação de Violência"
              color="#2196F3"
            />
            <CustomButton
              iconName="baby"
              label="Programa Mãe Coruja Pernambucana"
              color="#FF9800"
            />
            <CustomButton
              iconName="business"
              label="Secretaria da Mulher de Pernambuco"
              color="#D32F2F"
            />
            <CustomButton
              iconName="lightbulb"
              label="Bora Empreender Mulher"
              color="#FF6B6B"
            />
            <CustomButton
              iconName="shield"
              label="Patrulha Maria da Penha"
              color="#4CAF50"
            />
            <CustomButton
              iconName="megaphone"
              label="Ouvidoria da Mulher"
              color="#2196F3"
            />
          </VStack>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
