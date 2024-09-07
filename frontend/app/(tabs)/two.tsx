import { ScrollView, StyleSheet } from 'react-native';

import EditScreenInfo from '@/components/EditScreenInfo';
import { Text, View } from '@/components/Themed';
import { Card } from '@/components/ui/card';
import { Image } from '@/components/ui/image';
import { Heading } from '@/components/ui/heading';
import { Link, LinkText } from '@/components/ui/link';
import { HStack } from '@/components/ui/hstack';
import { ArrowRightIcon, Icon } from '@/components/ui/icon';


export default function TabTwoScreen() {
  return (
    <ScrollView>
    <Card className="p-5 rounded-lg max-w-[360px] m-3 mb-10">
    <Image
      
      source={{
        uri: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTv_3AHrrr0rHZ_SAMQrFX0Rp1SEioNfRqT1Q&s",
      }}
      className="mb-6 h-[240px] w-full rounded-md resize cover"
      
    />
    <Text className="text-sm font-normal mb-2 text-typography-700">
      Aug 28, 2024
    </Text>
    <Heading size="md" className="mb-4">
      Cresce o numero de denuncias utilizando o app PEM de forma anonima
    </Heading>
    <Link href="https://gluestack.io/" isExternal>
      <HStack className="items-center">
        <LinkText
          size="sm"
          className="font-semibold text-info-600 no-underline"
        >
          Saiba mais ...
        </LinkText>
        <Icon
          as={ArrowRightIcon}
          size="sm"
          className="text-info-600 mt-0.5 ml-0.5"
        />
      </HStack>
    </Link>
  </Card>

  <Card className="p-5 rounded-lg max-w-[360px] m-3 mb-10">
    <Image
      source={{
        uri: "https://www.gncnews.com.br/assets/uploads/cc6fdfcf32051fe32cd01666b6fc86f9.jpg",
      }}
      className="mb-6 h-[240px] w-full rounded-md"
    />
    <Text className="text-sm font-normal mb-2 text-typography-700">
      Jul 19, 2024
    </Text>
    <Heading size="md" className="mb-4">
      Aberto Inscrições para cursos de capacitação em beleza
    </Heading>
    <Link href="https://gluestack.io/" isExternal>
      <HStack className="items-center">
        <LinkText
          size="sm"
          className="font-semibold text-info-600 no-underline"
        >
          Saiba mais ...
        </LinkText>
        <Icon
          as={ArrowRightIcon}
          size="sm"
          className="text-info-600 mt-0.5 ml-0.5"
        />
      </HStack>
    </Link>
  </Card>
  </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  title: {
    fontSize: 20,
    fontWeight: 'bold',
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: '80%',
  },
});
