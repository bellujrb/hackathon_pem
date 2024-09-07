import React from "react";
import { SafeAreaView } from "react-native-safe-area-context";

import { useNavigation } from "expo-router";
import { VStack } from "./ui/vstack";
import { Text } from "./ui/text";
import { Box } from "./ui/box";
import { Button, ButtonText } from "./ui/button";
import { View } from "react-native";
import { Link } from "./ui/link";
import { Divider } from "./ui/divider";

export default function EditScreenInfo() {
  const navigation = useNavigation();

  return (
    <SafeAreaView
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        // backgroundColor: "#f5f5f5",
      }}
    >
      <View className="flex flex-1 items-center, justify-center">
        {/* Título */}
        <Text className="text-3xl text-center bold">
          A importância de utilizar este aplicativo
        </Text>
        <View className="mt-5 mb-5">
        <Divider />
        </View>
        

        {/* Mensagem de Apoio */}
        <Text className="text-center">
          Este aplicativo foi criado especialmente para oferecer suporte e
          proteção às mulheres. Todas as suas interações podem ser feitas de
          forma completamente anônima, sem a possibilidade de qualquer parceiro
          que já tenha atacado você descobrir suas ações. Seu bem-estar e
          segurança são nossa prioridade.
        </Text>

        {/* Mensagem de Segurança */}
        <Box bg="#ffeded" padding={16} borderRadius={10} mt={4}>
          <Text
            style={{
              fontSize: 14,
              color: "#ff6b6b",
              textAlign: "center",
              fontWeight: "bold",
            }}
          >
            Sua privacidade é garantida. Nenhum dado será compartilhado sem o
            seu consentimento.
          </Text>
        </Box>
      </View>
    </SafeAreaView>
  );
}
